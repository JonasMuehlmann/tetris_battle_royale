package ipc

import (
	"context"
	"fmt"
	"log"
	gameServiceProto "microservice/internal/core/protofiles/game_service"
	gameService "microservice/internal/core/services/game_service"
	"microservice/internal/core/types"
	"net"

	"google.golang.org/grpc"
)

//******************************************************************//
//                           GRPC related                           //
//******************************************************************//

type GameServiceServer struct {
	gameServiceProto.UnimplementedGameServiceServer
	GameService *gameService.GameService
	Logger      *log.Logger
}

func (service *GameServiceServer) StartGame(context context.Context, userIDList *gameServiceProto.UserIDList) (*gameServiceProto.EmptyMessage, error) {
	err := service.GameService.StartGame(userIDList.GetId())

	return &gameServiceProto.EmptyMessage{}, err
}

//******************************************************************//
//                              Adapter                            //
//******************************************************************//

type GameServiceIPCServerAdapter struct {
	GameServiceServer
	Logger *log.Logger
}

func (adapter *GameServiceIPCServerAdapter) Start(args interface{}) error {
	gameServiceArgs, ok := args.(types.DrivenAdapterGRPCArgs)
	if !ok {
		return fmt.Errorf("Invalid type %T for argument, expected %T", args, types.DrivenAdapterGRPCArgs{})
	}

	gameService_, ok := gameServiceArgs.Service.(*gameService.GameService)
	if !ok {
		var wanted *gameService.GameService
		return fmt.Errorf("Invalid type %T in argument %#v, expected %T", gameServiceArgs.Service, args, wanted)
	}
	// doesSatisfyPort := reflect.TypeOf(gameServiceArgs.Service).Implements(reflect.TypeOf((*drivingPorts.GameServicePort)(nil)).Elem())
	// if doesSatisfyPort {
	// 	var wanted *drivingPorts.GameServicePort
	// 	return fmt.Errorf("Invalid type %T in argument %#v, expected %T", gameServiceArgs.Service, args, wanted)
	// }

	listener := gameServiceArgs.Listener

	grpcServer := grpc.NewServer()
	gameServiceServer := &GameServiceServer{GameService: gameService_}
	// gameServiceServer := &GameServiceServer{GameService: (drivingPorts.GameServicePort)(gameServiceArgs.Service)}

	gameServiceProto.RegisterGameServiceServer(grpcServer, gameServiceServer)

	adapter.Logger.Printf("Starting GRPC server on port %v", listener.Addr().(*net.TCPAddr).Port)

	return grpcServer.Serve(listener)
}
