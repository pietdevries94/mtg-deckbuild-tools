package main

import (
	"github.com/labstack/echo"
	"github.com/pietdevries94/mtg-deckbuild-tools/backend/domains/entry"
)

func postEntry(c echo.Context) error {
	err := entry.AddCardEntryBySetAndNumber(c.FormValue(`set`), c.FormValue(`number`))
	if err != nil {
		return err
	}
	return c.String(201, ``)
}
