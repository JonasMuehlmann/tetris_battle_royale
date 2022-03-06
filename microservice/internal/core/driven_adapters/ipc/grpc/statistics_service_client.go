package ipc

import (
	"context"
	statisticsServiceProto "microservice/internal/core/protofiles/statistics_service"
	"microservice/internal/core/types"

	"google.golang.org/grpc"
)

type StatisticsServiceIPCClientAdapter struct {
	grpcClient statisticsServiceProto.StatisticsServiceClient
}

func (adapter StatisticsServiceIPCClientAdapter) AddMatchRecord(record types.MatchRecord) error {

	message := &statisticsServiceProto.MatchRecord{
		Id:           record.ID,
		UserId:       record.UserID,
		Win:          record.Win,
		Score:        int32(record.Score),
		Length:       int32(record.Length),
		Start:        record.Start.String(),
		RatingChange: int32(record.RatingChange),
	}

	_, err := adapter.grpcClient.AddMatchRecord(context.Background(), message)

	return err
}

func (adapter StatisticsServiceIPCClientAdapter) Start(args interface{}) error {
	grpcConn, err := grpc.Dial("statistics-service:8081", grpc.WithInsecure())
	if err != nil {
		return err
	}

	adapter.grpcClient = statisticsServiceProto.NewStatisticsServiceClient(grpcConn)

	return nil
}
