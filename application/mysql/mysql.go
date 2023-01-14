package mysql

import (
	"database/sql"
	"goober/database/mysql"
)

func DB() *sql.DB {
	return mysql.DB
}
