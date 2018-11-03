package main

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/pietdevries94/mtg-deckbuild-tools/backend/domains/card"
	"github.com/pietdevries94/mtg-deckbuild-tools/backend/domains/entry"
	"github.com/pietdevries94/mtg-deckbuild-tools/backend/domains/inventory"
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

func getCardByName(c echo.Context) error {
	name := c.Param(`name`)

	card, err := card.GetCardByName(name)
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

func deleteEntry(c echo.Context) error {
	scryfallID, listStringID := c.Param(`scryfall_id`), c.Param(`list_id`)

	listID, err := strconv.Atoi(listStringID)

	err = entry.DeleteCardEntry(scryfallID, listID)
	if err != nil {
		return err
	}
	return c.NoContent(201)
}

func deleteList(c echo.Context) error {
	listStringID := c.Param(`id`)

	listID, err := strconv.Atoi(listStringID)

	err = list.DeleteList(listID)
	if err != nil {
		return err
	}
	return c.NoContent(201)
}

func getListEntries(c echo.Context) error {
	listStringID := c.Param(`id`)

	listID, err := strconv.Atoi(listStringID)

	entries, err := entry.GetEntriesForList(listID)
	if err != nil {
		return err
	}

	return c.JSON(200, getListEntriesResponse{
		Entries: entries,
	})
}

func postInventory(c echo.Context) error {
	body := c.Request().Body
	if err := inventory.LoadInventory(body); err != nil {
		return err
	}
	return c.NoContent(203)
}

func putEntryBoard(c echo.Context) error {
	scryfallID, listStringID := c.Param(`scryfall_id`), c.Param(`list_id`)
	listID, err := strconv.Atoi(listStringID)
	if err != nil {
		return err
	}

	type data struct {
		Board string `json:"board"`
	}

	d := data{}
	c.Bind(&d)

	if err := entry.SetBoardForEntry(scryfallID, listID, d.Board); err != nil {
		return err
	}	

	return c.NoContent(200)
}
