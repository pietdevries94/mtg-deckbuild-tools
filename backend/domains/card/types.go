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
	CastingCost  string    `json:"casting_cost"`
	OnlinePrice  string    `json:"online_price"`
	CopiesOwned  int       `json:"copies_owned"`
	BuyLink      string    `json:"buy_link"`
	TypeLine     string    `json:"type_line"`
}
