package types

type PlayerRating struct {
	ID      int `db:"id" json:"id"`
	MMR     int `db:"mmr" json:"mmr"`
	KFactor int `db:"k_factor" json:"k_factor"`
}
