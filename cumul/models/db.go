package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// DB instance
var DB *sql.DB

// InitDB : initialize the db and connect
func InitDB() {
	database, err := sql.Open("mysql", "jiraiya:Shivi<323@tcp(database-1.caqh2nel7qhl.us-east-2.rds.amazonaws.com:3306)/cumul")

	if err != nil {
		panic(err.Error())
	}

	DB = database
}
