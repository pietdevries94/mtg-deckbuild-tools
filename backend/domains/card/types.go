package card

import (
	"time"
)

type Card struct {
	ScryfallID   string    `json:"scryfall_id"`
	SetCode      string    `json:"set_code"`
	SetNumber    string    `json:"set_number"`
	Name         string    `json:"name"`
	OracleID     string    `json:"oracle_id"`
	ThumbnailURL string    `json:"thumbnail_url"`
	UpdatedAt    time.Time `json:"updated_at"`
}
