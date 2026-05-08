package repository

import (
	"database/sql"
	"log"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "data/data.db")
	if err != nil {
		log.Fatal(err)
	}
	Migrate()
}

func Migrate() {
	CreateLessons()
	CreateGroups()
	CreateUsers()
	CreateUni()
	createSubGroups()
}
