package models

import "github.com/rogeriopiatek/goLang-todoAPI/db"

// Insert function will try to insert a todo in our DB
func Insert(todo Todo) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		//error will be treated on handler
		return
	}
	defer conn.Close()

	//insert into todos table, and returns the created ID
	sql := `INSERT INTO todos (title, description, done) VALUES ($1, $2, $3) RETURNING id`

	//executes the query passing the values needed, and scans the return to get (in this case) the ID
	err = conn.QueryRow(sql, todo.Title, todo.Description, todo.Done).Scan(&id)

	//implicit return, returns id and error like defined
	//handler will treat the error, and send a proper response to the user
	return
}
