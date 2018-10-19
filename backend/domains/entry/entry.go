package entry

import (
	"github.com/pietdevries94/mtg-deckbuild-tools/backend/domains/card"
)

func AddCardEntry(payload AddEntryPayload) error {
	c, err := card.GetCardByScryfallID(payload.ScryfallID)
	if err != nil {
		return err
	}

	return UpsertEntry(Entry{
		ScryfallID: c.ScryfallID,
	})
}
