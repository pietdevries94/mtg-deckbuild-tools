package inventory

import (
	"encoding/csv"
	"io"
	"strconv"

	"github.com/pietdevries94/mtg-deckbuild-tools/backend/domains/card"
	"github.com/pietdevries94/mtg-deckbuild-tools/backend/util"
)

func LoadInventory(reader io.ReadCloser) error {
	r := csv.NewReader(reader)

	records, err := r.ReadAll()
	if err != nil {
		return err
	}
	go parseRecords(records)
	return nil
}

func parseRecords(records [][]string) {
	for i, record := range records {
		if i == 0 {
			continue
		}
		amount, err := strconv.Atoi(record[0])
		util.PanicOnErr(err)
		c, err := card.GetCardByName(record[1])
		util.PanicOnErr(err)

		if c.CopiesOwned == amount {
			continue
		}

		err = SetCopiesOwned(c.ScryfallID, amount)
		util.PanicOnErr(err)
	}
}
