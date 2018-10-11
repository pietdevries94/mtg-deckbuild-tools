package main

import (
	"github.com/labstack/echo"
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
