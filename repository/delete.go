package repository

import (
	"github.com/me/activities/db"
	"github.com/me/activities/models"
)

func Delete(id uint, todo models.Todo) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	sql := `DELETE FROM todos WHERE id=$1`

	rst, err := conn.Exec(sql, id)
	if err != nil {
		return 0, err
	}
	
	rowsAffected, _ := rst.RowsAffected()

	return rowsAffected, nil
	
}