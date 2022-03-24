package types

type Scoreboard struct {
	PlayerPerformances []PlayerPerformance `json:"player_performances"`
	Winner             string
}
