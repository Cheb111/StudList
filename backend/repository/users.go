package repository

import (
	"stud/models"
)

func CreateUser(name string, gender string,subgroups_id int, password string, userRole string, groupId int, UniID int) (int, error) {

	result, err := db.Exec("INSERT INTO users (name, gender, password, uniID, group_id, role ) VALUES( ?, ?, ?, ?, ?, ?)",
		name, gender, password, UniID, groupId, subgroups_id, userRole)

	if err != nil {
		return 0, err
	}

	id, _ := result.LastInsertId()
	return int(id), nil

}

func GetLessonsByUser(groupId int, subgroups_id int) []models.Lesson {
	rows, err := db.Query(`
  SELECT id, title, day, time, description, groupid, subgroups_id
  FROM lessons
  WHERE groupid = ?
 `, groupId)

	if err != nil {
		return nil
	}
	defer rows.Close()

	var lessons []models.Lesson

	for rows.Next() {
		var l models.Lesson
		rows.Scan(&l.ID, &l.Title, &l.Day, &l.Time, &l.Description, &l.GroupId, &l.Subgroup_id)
		lessons = append(lessons, l)
	}

	return lessons
}
