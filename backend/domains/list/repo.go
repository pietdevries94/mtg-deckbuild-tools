package list

import (
	"database/sql"
	"fmt"

	"github.com/pietdevries94/mtg-deckbuild-tools/backend/data"
	"github.com/pietdevries94/mtg-deckbuild-tools/backend/util"
)

func getLists() ([]List, error) {
	db := data.GetDB()

	rows, err := db.Query(`select id, name from lists`)
	if err != nil {
		if err == sql.ErrNoRows {
			return []List{}, nil
		}
		return nil, err
	}

	lists := []List{}
	i := 0
	for rows.Next() {
		l := List{}
		err := rows.Scan(&l.ID, &l.Name)
		util.PanicOnErr(err)
		lists[i] = l
		i++
	}

	return nil, nil
}

func addList(name string) error {
	db := data.GetDB()

	checkQ := `select count(*) from lists where name = ?`
	row := db.QueryRow(checkQ, name)
	count := 0
	err := row.Scan(&count)
	util.PanicOnErr(err)

	if count != 0 {
		return fmt.Errorf("list '%s' already exists", name)
	}

	insertQ := `insert into lists(name) values(?)`
	_, err = db.Exec(insertQ, name)

	return err
}
