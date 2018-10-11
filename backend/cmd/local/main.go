package main

import (
	"github.com/labstack/echo"
	"github.com/pietdevries94/mtg-deckbuild-tools/backend/data"
)

func main() {
	data.Init()

	e := echo.New()
	e.POST("/entry", postEntry)
	e.Logger.Fatal(e.Start(":1323"))
}
