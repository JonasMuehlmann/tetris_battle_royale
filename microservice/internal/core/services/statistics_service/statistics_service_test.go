package statisticsService_test

import (
	"log"
	repository "microservice/internal/core/driven_adapters/repository/postgres"
	statisticsService "microservice/internal/core/services/statistics_service"
	"testing"

	"github.com/stretchr/testify/suite"
)

type statisticsServiceTestSuite struct {
	suite.Suite
	db      *repository.PostgresDatabase
	service statisticsService.StatisticsService
}

func (suite *statisticsServiceTestSuite) SetupTest() {
	suite.db = repository.MakeDefaultPostgresTestDB(log.Default())

	userRepo := repository.PostgresDatabaseUserRepository{Logger: log.Default(), PostgresDatabase: *suite.db}
	statisticsRepo := repository.PostgresDatabaseStatisticsRepository{Logger: log.Default(), PostgresDatabase: *suite.db}
	suite.service = statisticsService.StatisticsService{UserRepo: userRepo, StatisticsRepo: statisticsRepo, Logger: log.Default()}
}

func TestRunStatisticsServiceTestSuite(t *testing.T) {
	suite.Run(t, new(statisticsServiceTestSuite))
}

func (suite *statisticsServiceTestSuite) TestGetPlayerProfileBasic() {
	suite.Equal(1, 2, "Oh no!")
}
