package repository

import (
	"github.com/me/activities/db"
	"github.com/me/activities/models"
)

func Update(id uint, todo models.Todo) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	sql := `UPDATE todos SET title=$1, description=$2, done=$3 WHERE id=$4`

	rst, err := conn.Exec(sql, todo.Title, todo.Description, todo.Done, id)
	if err != nil {
		return 0, err
	}
	
	rowsAffected, _ := rst.RowsAffected()

	return rowsAffected, nil
	
}