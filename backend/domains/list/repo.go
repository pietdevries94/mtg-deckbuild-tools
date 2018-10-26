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
	for rows.Next() {
		l := List{}
		err := rows.Scan(&l.ID, &l.Name)
		util.PanicOnErr(err)
		lists = append(lists, l)
	}

	return lists, nil
}

func addList(name string) (int, error) {
	db := data.GetDB()

	checkQ := `select count(*) from lists where name = ?`
	row := db.QueryRow(checkQ, name)
	count := 0
	err := row.Scan(&count)
	util.PanicOnErr(err)

	if count != 0 {
		return -1, fmt.Errorf("list '%s' already exists", name)
	}

	insertQ := `insert into lists(name) values(?)`
	res, err := db.Exec(insertQ, name)
	if err != nil {
		return -1, err
	}

	id, err := res.LastInsertId()
	return int(id), err
}

func deleteListEntries(id int) error {
	db := data.GetDB()

	_, err := db.Exec(`delete from entries where list_id = ?`, id)
	return err
}

func deleteList(id int) error {
	db := data.GetDB()

	_, err := db.Exec(`delete from lists where id = ?`, id)
	return err
}

func getList(id int) (List, error) {
	return List{}, nil
}
