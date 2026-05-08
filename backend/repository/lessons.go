package repository

import (
	"log"
	"stud/models"
)

var Lessons []models.Lesson

func GetLessons() []models.Lesson {

	
	rows, err := db.Query("SELECT id, title, day, time, description, group_id, subgroup_id FROM lessons")
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()

	var lessons []models.Lesson

	for rows.Next() {
		var l models.Lesson
		err := rows.Scan(&l.ID, &l.Title, &l.Day, &l.Time, &l.Description, &l.GroupId, &l.Subgroup_id)
		if err != nil {
			log.Println(err)
			continue
		}
		lessons = append(lessons, l)
	}

	return lessons
}

func AddLesson(l models.Lesson) error {
  _, err := db.Exec(`
  INSERT INTO lessons (title, day, time, group_id, subgroup_id)
  VALUES (?, ?, ?, ?, ?)
  `,
    l.Title,
    l.Day,
    l.Time,
    l.GroupId,
    l.Subgroup_id,
  )

  return err
}

func DeleteLesson(id int) error {
	_, err := db.Exec("DELETE FROM lessons WHERE id = ?", id)
	return err
}
