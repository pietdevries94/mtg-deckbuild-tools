package data

import (
	scryfall "github.com/BlueMonday/go-scryfall"
	"github.com/pietdevries94/mtg-deckbuild-tools/backend/util"
)

var client *scryfall.Client

func initScryfall() {
	var err error
	client, err = scryfall.NewClient()
	util.PanicOnErr(err)
}
