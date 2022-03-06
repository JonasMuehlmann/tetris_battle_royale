package ipc

import (
	"context"
	"fmt"
	drivingPorts "microservice/internal/core/driving_ports"
	statisticsServiceProto "microservice/internal/core/protofiles/statistics_service"
	"microservice/internal/core/types"
	"time"

	"google.golang.org/grpc"
)

type StatisticsServiceIPCServerAdapter struct {
	StatisticsServiceServer
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

	return &statisticsServiceProto.EmptyRequest{}, (*service.statisticsService).AddMatchRecord(newRecord)
}

type StatisticsServiceServer struct {
	statisticsServiceProto.UnimplementedStatisticsServiceServer
	statisticsService *drivingPorts.StatisticsServicePort
}

func (adapter StatisticsServiceIPCServerAdapter) Start(args interface{}) error {
	statisticsServiceArgs, ok := args.(types.DrivenAdapterGRPCArgs)
	if !ok {
		return fmt.Errorf("Invalid type %T for argument, expected %T", args, types.DrivenAdapterGRPCArgs{})
	}

	statisticsService, ok := statisticsServiceArgs.Service.(*drivingPorts.StatisticsServicePort)
	if !ok {
		return fmt.Errorf("Invalid type %T in argument %+v, expected %T", statisticsServiceArgs.Service, args, types.DrivenAdapterGRPCArgs{}.Service)
	}

	listener := statisticsServiceArgs.Listener

	grpcServer := grpc.NewServer()
	statisticsServiceServer := &StatisticsServiceServer{statisticsService: statisticsService}

	statisticsServiceProto.RegisterStatisticsServiceServer(grpcServer, statisticsServiceServer)

	return grpcServer.Serve(listener)
}
