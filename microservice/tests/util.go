package tests

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func ResetDB(db *sqlx.DB) {
	_, err := db.Exec("TRUNCATE users, sessions, player_profiles, player_statistics,player_ratings, match_records CASCADE")
	if err != nil {
		log.Fatal(err)
	}
}
