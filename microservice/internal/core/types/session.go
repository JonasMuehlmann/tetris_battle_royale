package types

import "time"

type Session struct {
	ID           string    `db:"id"`
	UserID       string    `db:"user_id"`
	CreationTime time.Time `db:"creation_time"`
}
