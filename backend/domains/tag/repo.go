package tag

import (
	"database/sql"

	"github.com/pietdevries94/mtg-deckbuild-tools/backend/data"
	"github.com/pietdevries94/mtg-deckbuild-tools/backend/util"
)

func selectNames(q string, args ...interface{}) ([]string, error) {
	db := data.GetDB()

	rows, err := db.Query(`select name from tags`, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return []string{}, nil
		}
		return nil, err
	}

	tags := []string{}
	i := 0
	for rows.Next() {
		tag := ``
		err := rows.Scan(&tag)
		util.PanicOnErr(err)
		tags[i] = tag
		i++
	}

	return nil, nil
}

func getTags() ([]string, error) {
	return selectNames(`select name from tags`)
}

func getTagsForListID(listID int) ([]string, error) {
	return selectNames(`select t.name 
	from entries e 
	INNER join entries_tags et on e.id = et.entry_id 
	INNER join tags t on t.id = et.tag_id
	where e.list_id = ?`, listID)
}

func getTagsForEntryID(entryID int) ([]string, error) {
	return selectNames(`select t.name 
	FROM entries_tags et
	INNER join tags t on t.id = et.tag_id
	where et.entry_id = ?`, entryID)
}
