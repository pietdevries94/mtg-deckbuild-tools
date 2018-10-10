package data

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pietdevries94/mtg-deckbuild-tools/backend/util"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "file:live.db?cache=shared")
	util.PanicOnErr(err)

	migrate(db)
}
