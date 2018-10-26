package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/pietdevries94/mtg-deckbuild-tools/backend/data"
)

func main() {
	data.Init()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	api := e.Group("/api")

	api.GET("/card/set-number/:set/:number", getCardBySetAndNumber)
	api.GET("/card/name/:name", getCardByName)
	api.GET("/list", getLists)
	api.GET("/tag", getTags)

	api.POST("/entry", postEntry)
	api.POST("/list", postList)

	api.DELETE("/entry/:scryfall_id/:list_id", deleteEntry)

	fs := http.FileServer(data.GetFrontendFS())
	e.GET("*", echo.WrapHandler(fs))

	e.Logger.Fatal(e.Start(":1323"))
}
