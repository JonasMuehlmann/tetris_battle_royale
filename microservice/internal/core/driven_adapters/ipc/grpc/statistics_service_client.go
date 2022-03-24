package ipc

import (
	"context"
	"log"
	statisticsServiceProto "microservice/internal/core/protofiles/statistics_service"
	"microservice/internal/core/types"

	"google.golang.org/grpc"
)

type StatisticsServiceIPCClientAdapter struct {
	GrpcClient statisticsServiceProto.StatisticsServiceClient
	Logger     *log.Logger
}

func (adapter *StatisticsServiceIPCClientAdapter) AddMatchRecord(record types.MatchRecord) error {

	message := &statisticsServiceProto.MatchRecord{
		Id:           record.ID,
		UserId:       record.UserID,
		Win:          record.Win,
		Score:        int32(record.Score),
		Length:       int32(record.Length),
		Start:        record.Start.String(),
		RatingChange: int32(record.RatingChange),
	}

	_, err := adapter.GrpcClient.AddMatchRecord(context.Background(), message)

	return err
}

func (adapter *StatisticsServiceIPCClientAdapter) Start(args interface{}) error {
	serverAddr := "statistics-service:8081"

	grpcConn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		adapter.Logger.Printf("Error: %v\n", err)

		return err
	}

	adapter.Logger.Printf("Connected to GRPC server at %v", serverAddr)

	adapter.GrpcClient = statisticsServiceProto.NewStatisticsServiceClient(grpcConn)

	return nil
}
