package types

import "fmt"

type InvalidMatchIDError struct {
	MatchID string
}

func (err InvalidMatchIDError) Error() string {
	return fmt.Sprintf("Match %v does not exist ", err.MatchID)
}

type InvalidUserIDError struct {
	UserID string
}

func (err InvalidUserIDError) Error() string {
	return fmt.Sprintf("Match %v does not exist ", err.UserID)
}
