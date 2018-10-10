package data

import (
	"database/sql"

	scryfall "github.com/BlueMonday/go-scryfall"
)

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
