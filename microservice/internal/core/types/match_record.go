package types

import "time"

type MatchRecord struct {
	ID           int       `db:"id" json:"id"`
	UserID       string    `db:"user_id" json:"user_id"`
	Win          bool      `db:"win" json:"win"`
	Score        int       `db:"score" json:"score"`
	Start        time.Time `db:"start" json:"start"`
	Length       int       `db:"length" json:"length"`
	RatingChange int       `db:"rating_change" json:"rating_change"`
}
