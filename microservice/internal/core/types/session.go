package domain

import "time"

type Session struct {
	ID           int       `db:"id"`
	UserID       int       `db:"user_id"`
	CreationTime time.Time `db:"creation_time"`
}
