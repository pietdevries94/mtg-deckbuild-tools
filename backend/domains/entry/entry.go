package entry

import (
	"github.com/pietdevries94/mtg-deckbuild-tools/backend/domains/card"
)

func AddCardEntryBySetAndNumber(set, number string) error {
	c, err := card.GetCardBySetAndNumber(set, number)
	if err != nil {
		return err
	}

	return UpsertEntry(Entry{
		ScryfallID: c.ScryfallID,
	})
}
