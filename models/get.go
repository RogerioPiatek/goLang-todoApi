package models

import "github.com/rogeriopiatek/goLang-todoAPI/db"

// Returns an Todo based on the ID passed to the function
func Get(id int64) (todo Todo, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * from todos WHERE id=$1`, id)

	//one pointer for each Todo property
	err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

	return
}
