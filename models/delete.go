package models

import "github.com/rogeriopiatek/goLang-todoAPI/db"

// Updates an Todo based on the ID passed to this function
func Delete(id int64, todo Todo) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM todos WHERE id=$1`, id)
	if err != nil {
		return 0, err
	}

	//returns the number of rows affected (an Int)
	//useful on the handler, we validate if DELETE only changes one Row (because one ID was passed)
	return res.RowsAffected()
}
