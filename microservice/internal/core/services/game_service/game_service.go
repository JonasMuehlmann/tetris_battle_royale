package gameService

import (
	"log"
	drivenPorts "microservice/internal/core/driven_ports"
	repoPorts "microservice/internal/core/driven_ports/repository"
	types "microservice/internal/core/types"

	ipcPorts "microservice/internal/core/driven_ports/ipc"

	"time"

	"github.com/google/uuid"
)

type GameService struct {
	UserRepo repoPorts.UserRepositoryPort
	// This port/adapter might need refactoring
	Logger      *log.Logger
	Matches     map[string]Match
	IPCServer   ipcPorts.GameServiceIPCServerPort
	GameAdapter drivenPorts.GamePort
}

func MakeGameService(userRepo repoPorts.UserRepositoryPort, ipcServerAdapter ipcPorts.GameServiceIPCServerPort, gameAdapter drivenPorts.GamePort, logger *log.Logger) GameService {
	return GameService{
		UserRepo:    userRepo,
		Logger:      logger,
		Matches:     make(map[string]Match),
		IPCServer:   ipcServerAdapter,
		GameAdapter: gameAdapter,
	}
}

func (service GameService) StartGame(userIDList []string) error {
	matchID := uuid.NewString()

	players := map[string]Player{}
	for _, userID := range userIDList {
		// TODO: This should probably be refactored into a separate function and will include more complex setup logic
		players[userID] = Player{
			ID:        userID,
			Score:     0,
			Playfield: Playfield{},
		}

		// Build list of opponent user IDs
		opponentList := make([]types.Opponent, len(userIDList))
		opponentUserIDList := make([]string, len(userIDList))

		copy(opponentUserIDList, userIDList)

		for j, opponentUserID := range opponentUserIDList {
			if opponentUserID == userID {
				opponentUserIDList[j] = opponentUserIDList[len(opponentUserIDList)-1]
				opponentUserIDList = opponentUserIDList[:len(opponentUserIDList)-1]

				break
			}
		}

		// Build list of opponent user names
		for j, opponentUserID := range opponentUserIDList {
			user, err := service.UserRepo.GetUserFromID(opponentUserID)
			if err != nil {
				service.Logger.Printf("Error: %v\n", err)

				return err
			}

			opponentList = append(opponentList, types.Opponent{opponentUserIDList[j], user.Username})
		}

		err := service.GameAdapter.SendMatchStartNotice(userID, matchID, opponentList)
		if err != nil {
			service.Logger.Printf("Could not notify client %v of game start", userID)
			service.Logger.Printf("Error: %v\n", err)

			return err
		}
	}

	service.Matches[matchID] = Match{
		ID:      matchID,
		Players: players,
	}

	go service.StartGameInternal(matchID)

	return nil
}

func (service GameService) StartGameInternal(matchID string) error {
	time.Sleep(5)
	for _, v := range service.Matches[matchID].Players {
		v.Playfield.BlockPreview.MakeBlockPreview()
		v.Playfield.StartGame()

		err := service.GameAdapter.SendStartBlockPreview(v.ID, v.Playfield.BlockPreview)
		if err != nil {
			return err
		}
	}
	return nil
}

// NOTE: This function has nothing to do with the matchmaking
func (service GameService) ConnectPlayer(userID string, connection interface{}) error {
	return service.GameAdapter.ConnectPlayer(userID, connection)
}

func (service GameService) MoveBlock(userID string, matchID string, direction types.MoveDirection) error {

	success, player := service.validateUserAndMatch(userID, matchID)

	if !success {
		return nil
	}

	switch direction {
	case types.MoveLeft:
		player.Playfield.MoveBlockLeft()
	case types.MoveRight:
		player.Playfield.MoveBlockRight()
	case types.MoveDown:
		player.Playfield.MoveBlockDown()
	}

	return service.GameAdapter.SendUpdatedBlockState(userID, types.BlockState{
		BlockPosition:  player.Playfield.curBlockPosition,
		RotationChange: types.RotateNone,
	})
}

func (service GameService) RotateBlock(userID string, matchID string, direction types.RotationDirection) error {

	success, player := service.validateUserAndMatch(userID, matchID)
	if !success {
		return nil
	}

	switch direction {
	case types.RotateLeft:
		player.Playfield.RotateBlockClockwise()
	case types.RotateRight:
		player.Playfield.RotateBlockCounterClockwise()
	}

	return service.GameAdapter.SendUpdatedBlockState(userID, types.BlockState{
		BlockPosition:  player.Playfield.curBlockPosition,
		RotationChange: direction,
	})
}

func (service GameService) HardDropBlock(userID string, matchID string) error {
	success, player := service.validateUserAndMatch(userID, matchID)

	if !success {
		return nil
	}
	player.Playfield.HardDropBlock()

	return service.GameAdapter.SendUpdatedBlockState(userID, types.BlockState{
		BlockPosition:  player.Playfield.curBlockPosition,
		RotationChange: types.RotateNone,
	})
}

func (service GameService) ToggleSoftDrop(userID string, matchID string) error {
	success, player := service.validateUserAndMatch(userID, matchID)

	if !success {
		return nil
	}
	player.Playfield.ToggleSoftDrop()

	return service.GameAdapter.SendUpdatedBlockState(userID, types.BlockState{
		BlockPosition:  player.Playfield.curBlockPosition,
		RotationChange: types.RotationDirection("none"),
	})
}

func (service *GameService) validateUserAndMatch(userID string, matchID string) (bool, Player) {
	var player Player
	if _, ok := service.Matches[matchID]; !ok {
		service.Logger.Printf("The match %v does not exist.", matchID)
		return false, player
	}
	if _, ok := service.Matches[matchID].Players[userID]; !ok {
		service.Logger.Printf("The user is not a member of the match.")
		return false, player
	}
	return true, player
}
