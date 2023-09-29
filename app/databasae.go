package app

import (
	"courze-backend-app/helper"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:secretparah@/courze")
	helper.PanicIfError(err)

	return db
}
