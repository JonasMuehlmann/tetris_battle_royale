package gameService

import (
	"log"
	drivenPorts "microservice/internal/core/driven_ports"
	repoPorts "microservice/internal/core/driven_ports/repository"
	types "microservice/internal/core/types"

	ipcPorts "microservice/internal/core/driven_ports/ipc"

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

	players := [MatchSize]Player{}
	for i, userID := range userIDList {
		// TODO: This should probably be refactored into a separate function and will include more complex setup logic
		players[i] = Player{
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

	go service.Matches[matchID].Start()

	return nil
}

// NOTE: This function has nothing to do with the matchmaking
func (service GameService) ConnectPlayer(userID string, connection interface{}) error {
	return service.GameAdapter.ConnectPlayer(userID, connection)
}

func (service GameService) MoveBlock(userID string, matchID string, direction types.MoveDirection) error {

	if service.Matches[matchID] == nil {
		service.Logger.Printf("The match %v does not exist.", matchID)
		return nil
	} else if service.Matches[matchID].Players[userID] == nil {
		service.Logger.Printf("The user is not a member of the match.")
		return nil
	}

	switch direction {
	case types.MoveDirection.MoveLeft:
		service.Matches[matchID].Players[userID].Playfield.MoveBlockLeft()
		break
	case types.MoveDirection.MoveRight:
		service.Matches[matchID].Players[userID].Playfield.MoveBlockRight()
		break
	case types.MoveDirection.MoveDown:
		service.Matches[matchID].Players[userID].Playfield.MoveBlockDown()
		break
	}

	// TODO: Implement
	return nil
}

func (service GameService) RotateBlock(userID string, matchID string, direction types.RotationDirection) error {
	// TODO: Implement
	return nil
}

func (service GameService) HardDropBlock(userID string, matchID string) error {
	// TODO: Implement
	return nil
}

func (service GameService) ToggleSoftDrop(userID string, matchID string) error {
	// TODO: Implement
	return nil
}
