package data

import (
	"database/sql"
	"net/http"

	scryfall "github.com/BlueMonday/go-scryfall"
)

//go:generate go run assets_generate.go

func Init() {
	initDB()
	initScryfall()
}

func GetDB() *sql.DB {
	return db
}

func GetScryfall() *scryfall.Client {
	return client
}

func GetFrontendFS() http.FileSystem {
	return frontendFS
}
