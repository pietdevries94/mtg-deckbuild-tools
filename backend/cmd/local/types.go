package main

import (
	"github.com/pietdevries94/mtg-deckbuild-tools/backend/domains/card"
	"github.com/pietdevries94/mtg-deckbuild-tools/backend/domains/entry"
	"github.com/pietdevries94/mtg-deckbuild-tools/backend/domains/list"
)

type getCardResponse struct {
	Card    card.Card     `json:"card"`
	Entries []entry.Entry `json:"entries"`
}

type getListsResponse struct {
	Lists []list.List `json:"lists"`
}

type getTagsResponse struct {
	Tags []string `json:"tags"`
}
