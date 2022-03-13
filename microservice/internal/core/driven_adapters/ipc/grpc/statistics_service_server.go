package ipc

import (
	"context"
	"fmt"
	"log"
	statisticsServiceProto "microservice/internal/core/protofiles/statistics_service"
	statisticsService "microservice/internal/core/services/statistics_service"

	"microservice/internal/core/types"
	"strings"
	"time"

	"google.golang.org/grpc"
)

//******************************************************************//
//                           GRPC related                           //
//******************************************************************//

type StatisticsServiceServer struct {
	statisticsServiceProto.UnimplementedStatisticsServiceServer
	StatisticsService *statisticsService.StatisticsService
	Logger            *log.Logger
}

func (service StatisticsServiceServer) AddMatchRecord(context context.Context, record *statisticsServiceProto.MatchRecord) (*statisticsServiceProto.EmptyRequest, error) {
	startDateTime, err := time.Parse("2006-01-02 15:04:05", record.GetStart())
	if err != nil {
		return &statisticsServiceProto.EmptyRequest{}, err
	}

	newRecord := types.MatchRecord{
		ID:           record.GetId(),
		UserID:       record.GetUserId(),
		Win:          record.GetWin(),
		Score:        int(record.GetScore()),
		Start:        startDateTime,
		Length:       int(record.GetLength()),
		RatingChange: int(record.GetRatingChange()),
	}

	return &statisticsServiceProto.EmptyRequest{}, service.StatisticsService.AddMatchRecord(newRecord)
}

//******************************************************************//
//                              Adapter                            //
//******************************************************************//

type StatisticsServiceIPCServerAdapter struct {
	StatisticsServiceServer
	Logger *log.Logger
}

func (adapter StatisticsServiceIPCServerAdapter) Start(args interface{}) error {
	statisticsServiceArgs, ok := args.(types.DrivenAdapterGRPCArgs)
	if !ok {
		return fmt.Errorf("Invalid type %T for argument, expected %T", args, types.DrivenAdapterGRPCArgs{})
	}

	statisticsService_, ok := statisticsServiceArgs.Service.(*statisticsService.StatisticsService)
	if !ok {
		var wanted *statisticsService.StatisticsService
		return fmt.Errorf("Invalid type %T in argument %#v, expected %T", statisticsServiceArgs.Service, args, wanted)
	}
	// doesSatisfyPort := reflect.TypeOf(statisticsServiceArgs.Service).Implements(reflect.TypeOf((*drivingPorts.StatisticsServicePort)(nil)).Elem())
	// if doesSatisfyPort {
	// 	var wanted *drivingPorts.StatisticsServicePort
	// 	return fmt.Errorf("Invalid type %T in argument %#v, expected %T", statisticsServiceArgs.Service, args, wanted)
	// }

	listener := statisticsServiceArgs.Listener

	grpcServer := grpc.NewServer()
	statisticsServiceServer := &StatisticsServiceServer{StatisticsService: statisticsService_}
	// statisticsServiceServer := &StatisticsServiceServer{StatisticsService: (drivingPorts.StatisticsServicePort)(statisticsServiceArgs.Service)}

	statisticsServiceProto.RegisterStatisticsServiceServer(grpcServer, statisticsServiceServer)

	adapter.Logger.Printf("Starting GRPC server on port %v", strings.Split(listener.Addr().String(), ":")[1])

	return grpcServer.Serve(listener)
}
