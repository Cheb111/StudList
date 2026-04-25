package repository

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"stud/models"
)

var Lessons []models.Lesson

func GetLessons() []models.Lesson {
	rows, err := db.Query("SELECT id, title, day, time, description, group_id FROM lessons")
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()

	var lessons []models.Lesson

	for rows.Next() {
		var l models.Lesson
		err := rows.Scan(&l.ID, &l.Title, &l.Day, &l.Time, &l.Description, &l.GroupId)
		if err != nil {
			log.Println(err)
			continue
		}
		lessons = append(lessons, l)
	}

	return lessons
}

var db *sql.DB

func InitDB(filename string) {
	var err error
	db, err = sql.Open("sqlite3", filename)
	if err != nil {
		log.Fatal(err)
	}
	createTable()
}

func createTable() {
	lessontable := `
 CREATE TABLE IF NOT EXISTS lessons (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  title TEXT,
  day TEXT,
  time TEXT,
  description TEXT,
  group_id INTEGER
 );
 `

	_, err := db.Exec(lessontable)
	if err != nil {
		log.Fatal(err)
	}

	usertable := `
	CREATE TABLE IF NOT EXISTS users(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	groupid INTEGER,
	password TEXT,
	userRole TEXT
	);
	`

	_, err = db.Exec(usertable)

	if err != nil {
		log.Fatal(err)
	}
}

func CreateUser(name string, password string, userRole string) (int, error) {

	var groupid = 1

	result, err := db.Exec("INSERT INTO users (name, password, groupid, userRole) VALUES(?, ?, ?, ?)", name, password, groupid, userRole)

	if err != nil {
		return 0, err
	}

	id, _ := result.LastInsertId()
	return int(id), nil

}

//==============================================================

func AddLesson(l models.Lesson) error {
	query := `
 INSERT INTO lessons (title, day, time, description, group_id)
 VALUES (?, ?, ?, ?, ?)
 `

	_, err := db.Exec(query, l.Title, l.Day, l.Time, l.Description, l.GroupId)
	return err
}

func DeleteLesson(id int) error {
	_, err := db.Exec("DELETE FROM lessons WHERE id = ?", id)
	return err
}
