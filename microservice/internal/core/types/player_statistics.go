package types

type PlayerStatistics struct {
	ID             int     `db:"id" json:"id"`
	Score          int     `db:"score" json:"score"`
	ScorePerMinute float32 `db:"score_per_minute" json:"score_per_minute"`
	Wins           int     `db:"wins" json:"wins"`
	Losses         int     `db:"losses" json:"losses"`
	Winrate        float32 `db:"winrate" json:"winrate"`
	WinsAsTop10    int     `db:"wins_as_top_10" json:"wins_as_top_10"`
	WinsAsTop5     int     `db:"wins_as_top_5" json:"wins_as_top_5"`
	WinsAsTop3     int     `db:"wins_as_top_3" json:"wins_as_top_3"`
	WinsAsTop1     int     `db:"wins_as_top_1" json:"wins_as_top_1"`
}
