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

	defer db.Close()

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

func (suite *statisticsServiceTestSuite) TestGetPlayerStatisticsBasic() {
	db, err := suite.db.GetConnection()
	suite.NoError(err)

	defer db.Close()

	common.ResetDB(db)
	suite.NoError(err)

	_, err = db.Exec("INSERT INTO users(id) VALUES('123e4567-e89b-12d3-a456-426614174000')")
	suite.NoError(err)

	_, err = db.Exec("INSERT INTO player_ratings(id) VALUES(0)")
	suite.NoError(err)

	_, err = db.Exec("INSERT INTO player_statistics VALUES(0, 0, 0.0, 0, 0, 0.0, 0, 0, 0, 0)")
	suite.NoError(err)

	_, err = db.Exec("INSERT INTO player_profiles VALUES(0, '123e4567-e89b-12d3-a456-426614174000', 0, 0, 0, '1999-01-08 04:05:06')")
	suite.NoError(err)

	playerStatistics, err := suite.service.GetPlayerStatistics("123e4567-e89b-12d3-a456-426614174000")

	suite.Equal(playerStatistics, types.PlayerStatistics{
		ID:             0,
		Score:          0,
		ScorePerMinute: 0,
		Wins:           0,
		Losses:         0,
		Winrate:        0,
		WinsAsTop10:    0,
		WinsAsTop5:     0,
		WinsAsTop3:     0,
		WinsAsTop1:     0,
	})
	suite.NoError(err)
}

func (suite *statisticsServiceTestSuite) TestGetMatchRecordsBasic() {
	db, err := suite.db.GetConnection()
	suite.NoError(err)

	defer db.Close()

	common.ResetDB(db)
	suite.NoError(err)

	_, err = db.Exec("INSERT INTO users(id) VALUES('123e4567-e89b-12d3-a456-426614174000')")
	suite.NoError(err)

	_, err = db.Exec("INSERT INTO match_records VALUES('123e4567-e89b-12d3-a456-426614174000','123e4567-e89b-12d3-a456-426614174000', TRUE, 0, 0, '1999-01-08 04:05:06', 0)")
	suite.NoError(err)

	_, err = db.Exec("INSERT INTO match_records VALUES('123e4567-e89b-12d3-a456-426614173999','123e4567-e89b-12d3-a456-426614174000', TRUE, 0, 0, '1999-01-08 04:05:06', 0)")
	suite.NoError(err)

	matchRecords, err := suite.service.GetMatchRecords("123e4567-e89b-12d3-a456-426614174000")
	suite.NoError(err)
	suite.Equal(2, len(matchRecords))

	matchRecords[0].Start = matchRecords[0].Start.UTC()
	matchRecords[1].Start = matchRecords[1].Start.UTC()

	date, err := time.Parse("2006-01-02 15:04:05", "1999-01-08 04:05:06")
	suite.NoError(err)

	suite.Equal(matchRecords, []types.MatchRecord{
		{
			ID:           "123e4567-e89b-12d3-a456-426614174000",
			UserID:       "123e4567-e89b-12d3-a456-426614174000",
			Win:          true,
			Score:        0,
			Start:        date.UTC(),
			Length:       0,
			RatingChange: 0,
		},
		{
			ID:           "123e4567-e89b-12d3-a456-426614173999",
			UserID:       "123e4567-e89b-12d3-a456-426614174000",
			Win:          true,
			Score:        0,
			Start:        date.UTC(),
			Length:       0,
			RatingChange: 0,
		},
	})
}
func (suite *statisticsServiceTestSuite) TestGetMatchRecordBasic() {
	db, err := suite.db.GetConnection()
	suite.NoError(err)

	defer db.Close()

	common.ResetDB(db)
	suite.NoError(err)

	_, err = db.Exec("INSERT INTO users(id) VALUES('123e4567-e89b-12d3-a456-426614174000')")
	suite.NoError(err)

	_, err = db.Exec("INSERT INTO match_records VALUES('123e4567-e89b-12d3-a456-426614174000','123e4567-e89b-12d3-a456-426614174000', TRUE, 0, 0, '1999-01-08 04:05:06', 0)")
	suite.NoError(err)

	matchRecord, err := suite.service.GetMatchRecord("123e4567-e89b-12d3-a456-426614174000")
	suite.NoError(err)

	matchRecord.Start = matchRecord.Start.UTC()

	date, err := time.Parse("2006-01-02 15:04:05", "1999-01-08 04:05:06")
	suite.NoError(err)

	suite.Equal(matchRecord, types.MatchRecord{
		ID:           "123e4567-e89b-12d3-a456-426614174000",
		UserID:       "123e4567-e89b-12d3-a456-426614174000",
		Win:          true,
		Score:        0,
		Start:        date.UTC(),
		Length:       0,
		RatingChange: 0,
	})
}
