package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/pietdevries94/mtg-deckbuild-tools/backend/data"
)

func main() {
	data.Init()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	e.GET("/card/set-number/:set/:number", getCardBySetAndNumber)
	e.POST("/entry", postEntry)

	e.Logger.Fatal(e.Start(":1323"))
}
