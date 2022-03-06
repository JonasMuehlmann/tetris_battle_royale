package types

import "time"

type PlayerProfile struct {
	ID                 int       `db:"id" json:"id"`
	UserID             string    `db:"user_id" json:"user_id"`
	Playtime           int       `db:"playtime" json:"playtime"`
	PlayerRatingID     int       `db:"player_rating_id" json:"player_rating_id"`
	PlayerStatisticsID int       `db:"player_statistics_id" json:"player_statistics_id"`
	LastUpdate         time.Time `json:"last_update" db:"last_update"`
}
