package entry

import (
	"github.com/pietdevries94/mtg-deckbuild-tools/backend/data"
)

func UpsertEntry(entry Entry) error {
	db := data.GetDB()

	q := `insert or replace into entries (oracle_id) values (?)`
	_, err := db.Exec(q, entry.OracleID)

	return err
}
