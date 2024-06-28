package models

import "github.com/rogeriopiatek/goLang-todoAPI/db"

// Updates an Todo based on the ID passed to this function
func Update(id int64, todo Todo) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE todos SET title=$2, description=$3, done=$4 WHERE id=$1`, id, todo.Title, todo.Description, todo.Done)
	if err != nil {
		return 0, err
	}

	//returns the number of rows affected (an Int)
	//useful on the handler, we validate if UPDATE only changes one Row (because one ID was passed)
	return res.RowsAffected()
}
