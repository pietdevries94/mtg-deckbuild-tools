package entry

import (
	"github.com/pietdevries94/mtg-deckbuild-tools/backend/domains/card"
)

type Entry struct {
	ID         int       `json:"id"`
	ScryfallID string    `json:"scryfall_id"`
	ListID     int       `json:"list_id"`
	Tags       []string  `json:"tags"`
	Card       card.Card `json:"card"`
}

type AddEntryPayload struct {
	ScryfallID string   `json:"scryfall_id"`
	ListID     int      `json:"list_id"`
	Tags       []string `json:"tags"`
}
