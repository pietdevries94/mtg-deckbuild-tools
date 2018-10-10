package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/pietdevries94/mtg-deckbuild-tools/backend/data/database"
)

func main() {
	database.Init()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
