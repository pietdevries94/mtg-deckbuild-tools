package data

import (
	"database/sql"
	"fmt"

	"github.com/pietdevries94/mtg-deckbuild-tools/backend/util"
)

func migrate(db *sql.DB) {
	row := db.QueryRow(`SELECT 1 FROM sqlite_master WHERE type='table' AND name='migrations';`)

	hasTable := false
	if err := row.Scan(&hasTable); err != nil {
		if err != sql.ErrNoRows {
			panic(err)
		}
	}

	if !hasTable {
		_, err := db.Exec(`CREATE TABLE migrations(ID INTEGER PRIMARY KEY ASC)`)
		util.PanicOnErr(err)
	}

	row = db.QueryRow(`SELECT coalesce(MAX(ID), -1) FROM migrations`)

	lastID := 0
	err := row.Scan(&lastID)
	util.PanicOnErr(err)

	migrations := getMigrations()
	if lastID >= 0 && len(migrations) > lastID {
		migrations = migrations[lastID+1:]
	}
	for id, migration := range migrations {
		_, err := db.Exec(migration)
		util.PanicOnErr(err)
		_, err = db.Exec(fmt.Sprintf(`INSERT INTO migrations VALUES (%d)`, id))
		util.PanicOnErr(err)
	}
}

func getMigrations() []string {
	return []string{
		`CREATE TABLE cards(
			scryfall_id TEXT PRIMARY KEY,
			set_code TEXT,
			set_number TEXT,
			name TEXT,
			oracle_id TEXT,
			updated_at DATETIME
		)`,
		`CREATE TABLE entries(
			scryfall_id TEXT,
			FOREIGN KEY(scryfall_id) REFERENCES cards(scryfall_id)
		)`,
		`ALTER TABLE cards ADD thumbnail_url TEXT`,
	}
}
