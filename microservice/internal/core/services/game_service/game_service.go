package gameService

import (
	"fmt"
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
	Logger              *log.Logger
	Matches             map[string]Match
	IPCServer           ipcPorts.GameServiceIPCServerPort
	GameAdapter         drivenPorts.GamePort
	StatisticsIPCClient ipcPorts.StatisticsServiceIPCClientPort
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

func (service *GameService) StartGame(userIDList []string) error {
	matchID := uuid.NewString()

	players := map[string]Player{}
	for _, userID := range userIDList {
		player, err := service.initPlayer(players, userID, userIDList, matchID)
		if err != nil {
			log.Printf("Error: %v\n", err)

			return err
		}

		players[userID] = player
	}

	service.Matches[matchID] = Match{
		ID:                 matchID,
		Players:            players,
		PlayerEliminations: make(chan string, 10),
	}

	go service.StartGameInternal(matchID)

	return nil
}

func (service *GameService) StartGameInternal(matchID string) error {
	time.Sleep(5)
	for _, v := range service.Matches[matchID].Players {
		v.Playfield.TetrominoPreview = MakeTetrominoPreview()
		v.Playfield.StartGame()

		var tetrominos []types.Tetromino
		for e := v.Playfield.TetrominoPreview.tetrominoQueue.Front(); e != nil; e = e.Next() {
			tetrominos = append(tetrominos, types.Tetromino(e.Value.(types.Tetromino)))
		}

		err := service.GameAdapter.SendStartTetrominoPreview(v.ID, tetrominos)
		if err != nil {
			return err
		}
	}
	return nil
}

func (service GameService) StopGame(matchID string) error {
	return nil
}

// NOTE: This function has nothing to do with the matchmaking
func (service *GameService) ConnectPlayer(userID string, connection interface{}) error {
	return service.GameAdapter.ConnectPlayer(userID, connection)
}

func (service *GameService) MoveTetromino(userID string, matchID string, direction types.MoveDirection) error {

	err, player := service.validateUserAndMatch(userID, matchID)

	if err != nil {
		service.Logger.Printf("Error: %v\n", err)

		return err
	}

	switch direction {
	case types.MoveLeft:
		player.Playfield.MoveTetrominoLeft()
	case types.MoveRight:
		player.Playfield.MoveTetrominoRight()
	case types.MoveDown:
		player.Playfield.MoveTetrominoDown()
	}

	return service.GameAdapter.SendUpdatedTetrominoState(userID, types.TetrominoState{
		TetrominoPosition: player.Playfield.curTetrominoPosition,
		RotationChange:    types.RotateNone,
	})
}

func (service *GameService) RotateTetromino(userID string, matchID string, direction types.RotationDirection) error {

	err, player := service.validateUserAndMatch(userID, matchID)

	if err != nil {
		service.Logger.Printf("Error: %v\n", err)

		return err
	}

	switch direction {
	case types.RotateLeft:
		player.Playfield.RotateTetrominoClockwise()
	case types.RotateRight:
		player.Playfield.RotateTetrominoCounterClockwise()
	default:
		return fmt.Errorf(`Received invalid rotation direction "%v"`, direction)
	}

	return service.GameAdapter.SendUpdatedTetrominoState(userID, types.TetrominoState{
		TetrominoPosition: player.Playfield.curTetrominoPosition,
		RotationChange:    direction,
	})
}

func (service GameService) HardDropTetromino(userID string, matchID string) error {
	err, player := service.validateUserAndMatch(userID, matchID)

	if err != nil {
		service.Logger.Printf("Error: %v\n", err)

		return err
	}

	player.Playfield.HardDropTetromino()

	return service.GameAdapter.SendUpdatedTetrominoState(userID, types.TetrominoState{
		TetrominoPosition: player.Playfield.curTetrominoPosition,
		RotationChange:    types.RotateNone,
	})
}

func (service *GameService) ToggleSoftDrop(userID string, matchID string) error {
	err, player := service.validateUserAndMatch(userID, matchID)

	if err != nil {
		service.Logger.Printf("Error: %v\n", err)

		return err
	}

	player.Playfield.ToggleSoftDrop()

	return service.GameAdapter.SendUpdatedTetrominoState(userID, types.TetrominoState{
		TetrominoPosition: player.Playfield.curTetrominoPosition,
		RotationChange:    types.RotateNone,
	})
}

func (service *GameService) validateUserAndMatch(userID string, matchID string) (error, Player) {
	match, ok := service.Matches[matchID]
	if !ok {
		err := types.InvalidMatchIDError{matchID}

		return err, Player{}
	}

	player, ok := match.Players[userID]
	if !ok {
		err := types.InvalidUserIDError{userID}

		return err, Player{}
	}

	return nil, player
}

func (service *GameService) buildOpponentList(userIDList []string, userID string) ([]types.Opponent, error) {
	opponentList := make([]types.Opponent, len(userIDList))
	opponentUserIDList := make([]string, len(userIDList))

	copy(opponentUserIDList, userIDList)

	// Build list of opponent user IDs
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

			return nil, err
		}

		opponentList = append(opponentList, types.Opponent{opponentUserIDList[j], user.Username})
	}
	return opponentList, nil
}

func (service *GameService) initPlayer(players map[string]Player, userID string, userIDList []string, matchID string) (Player, error) {
	player := Player{
		ID:        userID,
		Score:     0,
		Playfield: Playfield{},
	}

	opponentList, err := service.buildOpponentList(userIDList, userID)
	if err != nil {
		log.Printf("Error: %v\n", err)

		return Player{}, err
	}

	err = service.GameAdapter.SendMatchStartNotice(userID, matchID, opponentList)
	if err != nil {
		service.Logger.Printf("Could not notify client %v of game start", userID)
		service.Logger.Printf("Error: %v\n", err)

		return Player{}, err
	}
	return player, nil
}
