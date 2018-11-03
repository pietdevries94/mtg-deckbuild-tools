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
	start := 0
	if lastID >= 0 && len(migrations) > lastID {
		start = lastID + 1
	}

	for i := start; i < len(migrations); i++ {
		_, err := db.Exec(migrations[i])
		util.PanicOnErr(err)
		fmt.Printf("migration %d\n", i)
		_, err = db.Exec(`INSERT INTO migrations(ID) VALUES (?)`, i)
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
			thumbnail_url TEXT,
			updated_at DATETIME
		)`,
		`CREATE TABLE lists(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT
		)`,
		`CREATE TABLE entries(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			scryfall_id TEXT,
			list_id INTEGER,
			FOREIGN KEY(scryfall_id) REFERENCES cards(scryfall_id),
			FOREIGN KEY(list_id) REFERENCES lists(id)
		)`,
		`CREATE TABLE tags(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT
		)`,
		`CREATE TABLE entries_tags(
			entry_id INTEGER,
			tag_id INTEGER,
			FOREIGN KEY(entry_id) REFERENCES entries(id),
			FOREIGN KEY(tag_id) REFERENCES tag(id)
		)`,
		`ALTER TABLE cards ADD COLUMN casting_cost TEXT`,
		`ALTER TABLE cards ADD COLUMN online_price TEXT`,
		`ALTER TABLE cards ADD COLUMN copies_owned INTEGER DEFAULT 0`,
		`ALTER TABLE cards ADD COLUMN buy_link TEXT`,
		`ALTER TABLE cards ADD COLUMN type_line TEXT`,
		`ALTER TABLE entries ADD COLUMN board TEXT DEFAULT ''`,
	}
}
