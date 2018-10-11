package entry

import (
	"github.com/pietdevries94/mtg-deckbuild-tools/backend/data"
)

func UpsertEntry(entry Entry) error {
	db := data.GetDB()

	q := `insert or replace into entries (scryfall_id) values (?)`
	_, err := db.Exec(q, entry.ScryfallID)

	return err
}
