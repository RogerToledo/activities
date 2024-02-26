package repository

import (
	"github.com/me/activities/db"
	"github.com/me/activities/models"
)

func ReadAll() (todos []models.Todo, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `SELECT * FROM todo`

	rows, err := conn.Query(sql)
	if err != nil {
		return
	}

	for rows.Next() {
		var todo models.Todo

		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
		if err != nil {
			continue
			// TODO log error
		}

		todos = append(todos, todo)
	}

	return
}
