package entry

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/pietdevries94/mtg-deckbuild-tools/backend/data"
	"github.com/pietdevries94/mtg-deckbuild-tools/backend/util"
)

func insertEntry(entry Entry) (int, error) {
	db := data.GetDB()

	q := `insert into entries (scryfall_id, list_id) values (?, ?)`
	res, err := db.Exec(q, entry.ScryfallID, entry.ListID)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	return int(id), err
}

func getEntryID(entry Entry) (int, bool) {
	db := data.GetDB()

	q := `select id from entries where scryfall_id = ? and list_id = ?`
	row := db.QueryRow(q, entry.ScryfallID, entry.ListID)
	id := 0
	err := row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, false
		}
		panic(err)
	}

	return id, true
}

func getEntriesForCard(scryfallID string) ([]Entry, error) {
	db := data.GetDB()

	q := `select id, list_id, scryfall_id from entries where scryfall_id = ?`
	rows, err := db.Query(q, scryfallID)
	if err != nil {
		if err == sql.ErrNoRows {
			return []Entry{}, nil
		}
		return nil, err
	}

	entries := []Entry{}
	i := 0
	for rows.Next() {
		e := Entry{}
		err := rows.Scan(&e.ID, &e.ListID, &e.ScryfallID)
		util.PanicOnErr(err)
		entries[i] = e
		i++
	}

	return entries, nil
}

func addTags(entryID int, tags []string) error {
	db := data.GetDB()

	// Get the current rows
	currentRows, err := db.Query("select tag_id from entries_tags where entry_id = ?", entryID)
	util.PanicOnErr(err)

	existingTagIDs := []int{}
	i := 0
	for currentRows.Next() {
		id := 0
		currentRows.Scan(&id)
		existingTagIDs[i] = id
		i++
	}

	// Get all needed tag ids
	tagsString := strings.Join(tags, `','`)
	getIDsQuery := fmt.Sprintf("select id from tags where name in ('%s')", tagsString)
	rows, err := db.Query(getIDsQuery)
	util.PanicOnErr(err)

	// prepare insert
	insertQuery := fmt.Sprintf("insert into entries_tags(entry_id, tag_id) values(%d, ?)", entryID)
	prep, err := db.Prepare(insertQuery)
	util.PanicOnErr(err)

	for rows.Next() {
		id := 0
		err := rows.Scan(&id)
		util.PanicOnErr(err)

		// skip if already exists and remove the tag id from the list
		if index, ok := util.IntInSlice(id, existingTagIDs); ok {
			existingTagIDs = util.DeleteFromIntSlice(existingTagIDs, index)
			continue
		}

		_, err = prep.Exec(id)
		if err != nil {
			return err
		}
	}

	// delete unused existing
	tagIDsString := util.IntSliceJoin(existingTagIDs, `,`)
	deleteQuery := fmt.Sprintf("delete from entries_tags where tag_id in (%s) and entry_id = ?", tagIDsString)
	_, err = db.Exec(deleteQuery, entryID)
	util.PanicOnErr(err)

	return nil
}
