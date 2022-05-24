package models

import "github.com/stephan-lopes/to-do/db"

func Update(id int64, todo Todo) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	result, err := conn.Exec(`UPDATE todos SET title = $1, description = $2, completed = $3 WHERE id = $4`, todo.Title, todo.Description, todo.Completed, id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
