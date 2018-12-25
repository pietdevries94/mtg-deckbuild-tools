package inventory

import (
	"encoding/csv"
	"fmt"
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
		fmt.Printf("Parsing: %s\n", record[1])
		amount, err := strconv.Atoi(record[0])
		util.PanicOnErr(err)
		c, err := card.GetCardByName(record[1])
		util.PanicOnErr(err)

		fmt.Printf("Old: %d | New: %d\n", c.CopiesOwned, amount)

		if c.CopiesOwned == amount {
			continue
		}

		err = SetCopiesOwned(c.ScryfallID, amount)
		util.PanicOnErr(err)

		fmt.Printf("Updated %s\n", record[1])
	}
}
