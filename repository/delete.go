package repository

import (
	"github.com/me/todo-api/db"
)

func Delete(id uint) (int64, error) {
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
