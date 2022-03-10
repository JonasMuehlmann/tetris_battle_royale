package ipc

import (
	"context"
	gameServiceProto "microservice/internal/core/protofiles/game_service"
	gameService "microservice/internal/core/services/game_service"
)

type GameServiceServer struct {
	gameServiceProto.UnimplementedGameServiceServer
	GameService gameService.GameService
}

func (service GameServiceServer) StartGame(context context.Context, userIDList *gameServiceProto.UserIDList) (*gameServiceProto.EmptyMessage, error) {
	err := service.GameService.StartGame(userIDList.GetId())
	if err != nil {
		return &gameServiceProto.EmptyMessage{}, err
	}

	return &gameServiceProto.EmptyMessage{}, nil
}
