package entry

import (
	"errors"

	"github.com/pietdevries94/mtg-deckbuild-tools/backend/domains/card"
)

var (
	ErrEntryNotFound = errors.New(`entry not found`)
)

type Entry struct {
	ID         int       `json:"id"`
	ScryfallID string    `json:"scryfall_id"`
	ListID     int       `json:"list_id"`
	Tags       []string  `json:"tags"`
	Card       card.Card `json:"card"`
	Board      string    `json:"board"`
}

type AddEntryPayload struct {
	ScryfallID string   `json:"scryfall_id"`
	ListID     int      `json:"list_id"`
	Tags       []string `json:"tags"`
}

type SetBoardPayload struct {
	ScryfallID string `json:"scryfall_id"`
	ListID     int    `json:"list_id"`
	Board      string `json:"board"`
}
