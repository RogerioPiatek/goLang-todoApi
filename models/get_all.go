package models

import "github.com/rogeriopiatek/goLang-todoAPI/db"

// Get all Todo's from the DB
func GetAll() (todos []Todo, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * from todos`)
	if err != nil {
		return
	}

	//loop to iterate over each row, and assign it properly
	for rows.Next() {
		//one Todo
		var todo Todo

		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
		if err != nil {
			continue
		}

		todos = append(todos, todo)
	}
	return
}
