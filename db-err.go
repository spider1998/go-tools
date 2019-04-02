package routing

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

func IsDBDuplicatedErr(err error) bool {
	if dbErr, ok := err.(*mysql.MySQLError); ok {
		if dbErr.Number == 1062 {
			return true
		}
	}
	return false
}

func IsDBNotFound(err error) bool {
	return err == sql.ErrNoRows
}


