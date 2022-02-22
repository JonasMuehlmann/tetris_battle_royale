package domain

import "time"

type Session struct {
	ID           int       `db:"ID"`
	UserID       int       `db:"user_ID"`
	CreationTime time.Time `db:"creation_time"`
}
