package inventory

import (
	"github.com/pietdevries94/mtg-deckbuild-tools/backend/data"
)

func SetCopiesOwned(scryfallID string, amount int) error {
	var db = data.GetDB()

	q := `update cards set copies_owned = ? where scryfall_id = ?`
	_, err := db.Exec(q, amount, scryfallID)

	return err
}
