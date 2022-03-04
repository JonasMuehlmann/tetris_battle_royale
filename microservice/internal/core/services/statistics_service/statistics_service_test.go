package statisticsService_test

import (
	"log"
	common "microservice/internal"
	repository "microservice/internal/core/driven_adapters/repository/postgres"
	statisticsService "microservice/internal/core/services/statistics_service"
	"microservice/internal/core/types"
	"testing"
	"time"

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
	db, err := suite.db.GetConnection()
	suite.NoError(err)

	common.ResetDB(db)

	_, err = db.Exec("INSERT INTO users(id) VALUES('123e4567-e89b-12d3-a456-426614174000')")
	suite.NoError(err)

	_, err = db.Exec("INSERT INTO player_ratings(id) VALUES(0)")
	suite.NoError(err)

	_, err = db.Exec("INSERT INTO player_statistics(id) VALUES(0)")
	suite.NoError(err)

	_, err = db.Exec("INSERT INTO player_profiles VALUES(0, '123e4567-e89b-12d3-a456-426614174000', 0, 0, 0, '1999-01-08 04:05:06')")
	suite.NoError(err)

	playerProfile, err := suite.service.GetPlayerProfile("123e4567-e89b-12d3-a456-426614174000")
	playerProfile.LastUpdate = playerProfile.LastUpdate.UTC()

	date, err := time.Parse("2006-01-02 15:04:05", "1999-01-08 04:05:06")
	suite.NoError(err)
	suite.Equal(playerProfile, types.PlayerProfile{
		ID:                 0,
		UserID:             "123e4567-e89b-12d3-a456-426614174000",
		Playtime:           0,
		PlayerRatingID:     0,
		PlayerStatisticsID: 0,
		LastUpdate:         date.UTC(),
	})
	suite.NoError(err)
}
