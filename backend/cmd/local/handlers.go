package main

import (
	"github.com/labstack/echo"
	"github.com/pietdevries94/mtg-deckbuild-tools/backend/domains/card"
	"github.com/pietdevries94/mtg-deckbuild-tools/backend/domains/entry"
	"github.com/pietdevries94/mtg-deckbuild-tools/backend/domains/list"
	"github.com/pietdevries94/mtg-deckbuild-tools/backend/domains/tag"
)

func postEntry(c echo.Context) error {
	payload := entry.AddEntryPayload{}
	c.Bind(&payload)

	err := entry.AddCardEntry(payload)
	if err != nil {
		return err
	}
	return c.String(201, ``)
}

func postList(c echo.Context) error {
	payload := list.AddListPayload{}
	c.Bind(&payload)

	id, err := list.AddList(payload)
	if err != nil {
		return err
	}
	return c.JSON(200, genericInsertResponse{
		ID: id,
	})
}

func getCardBySetAndNumber(c echo.Context) error {
	set, number := c.Param(`set`), c.Param(`number`)

	card, err := card.GetCardBySetAndNumber(set, number)
	if err != nil {
		return err
	}
	entries, err := entry.GetEntriesForCard(card.ScryfallID)
	if err != nil {
		return err
	}

	return c.JSON(200, getCardResponse{
		Card:    card,
		Entries: entries,
	})
}

func getLists(c echo.Context) error {
	lists, err := list.GetLists()
	if err != nil {
		return err
	}

	return c.JSON(200, getListsResponse{
		Lists: lists,
	})
}

func getTags(c echo.Context) error {
	tags, err := tag.GetTags()
	if err != nil {
		return err
	}

	return c.JSON(200, getTagsResponse{
		Tags: tags,
	})
}
