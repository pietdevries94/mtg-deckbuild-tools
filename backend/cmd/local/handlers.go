package main

import (
	"github.com/labstack/echo"
	"github.com/pietdevries94/mtg-deckbuild-tools/backend/domains/card"
	"github.com/pietdevries94/mtg-deckbuild-tools/backend/domains/entry"
)

type postEntryData struct {
	Set    string
	Number string
}

func postEntry(c echo.Context) error {
	d := postEntryData{}
	c.Bind(&d)

	err := entry.AddCardEntryBySetAndNumber(d.Set, d.Number)
	if err != nil {
		return err
	}
	return c.String(201, ``)
}

func getCardBySetAndNumber(c echo.Context) error {
	set, number := c.Param(`set`), c.Param(`number`)

	card, err := card.GetCardBySetAndNumber(set, number)
	if err != nil {
		return err
	}
	return c.JSON(200, card)
}
