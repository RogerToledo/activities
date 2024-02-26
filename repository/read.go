package repository

import (
	"github.com/me/activities/db"
	"github.com/me/activities/models"
)

func Read(id uint) (todo models.Todo, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `SELECT * FROM todo WHERE id = $1`

	row := conn.QueryRow(sql, id)

	err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

	return
}
