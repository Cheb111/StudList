package repository

import (
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func CreateLessons() {
	lessontable := `
 CREATE TABLE IF NOT EXISTS lessons (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  title TEXT,
  day TEXT,
  time TEXT,
  description TEXT,
  group_id INTEGER,
  subgroup_id INTEGER,
  user TEXT 
 
 );
 `

	_, err := db.Exec(lessontable)

	if err != nil {
		fmt.Println("Error 1", err)
		log.Fatal(err)

	}
}

func CreateUsers() {

	usertable := `
	CREATE TABLE IF NOT EXISTS users(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	gender TEXT,
	subgroup INTEGER,
	group_id INTEGER,
	password TEXT,
	role TEXT,
	uniID INTEGER
	);
	`

	_, err := db.Exec(usertable)

	if err != nil {
		fmt.Println("Error 2", err)
		log.Fatal(err)

	}
}

func CreateGroups() {
	grouptable := `
CREATE TABLE IF NOT EXISTS groups (
  	id INTEGER PRIMARY KEY AUTOINCREMENT,
  	name TEXT,
  	course INTEGER,
  	university_id INTEGER,
	UNIQUE(name, course, university_id)
);
`
	_, err := db.Exec(grouptable)

	if err != nil {
		fmt.Println("Error 3", err)
		log.Fatal(err)
	}

}

func CreateUni() {

	uniTable := `
CREATE TABLE IF NOT EXISTS universities (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT UNIQUE
);
`

	_, err := db.Exec(uniTable)

	if err != nil {
		log.Fatal(err)
	}
}

func createSubGroups() {

	SubGroupTable := `
CREATE TABLE IF NOT EXISTS subgroups (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT,
  group_id INTEGER
);
`

	_, err := db.Exec(SubGroupTable)

	if err != nil {
		log.Fatal(err)
	}
}
