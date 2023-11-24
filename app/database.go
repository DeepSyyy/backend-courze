package app

import (
	"courze-backend-app/helper"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:Falah190303@tcp(localhost:3306)/courze")
	helper.PanicIfError(err)

	return db
}
