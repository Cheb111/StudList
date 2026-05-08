package repository

import "stud/models"

func CreateSubGroup(name string, GroupID int) error {

	_, err := db.Exec(
		"INSERT INTO subgroups (name, group_id) VALUE (?, ?)",
		name, GroupID,
	)
	return err
}

func GetSubgroups(groupID int, userSubGroup_id int) ([]models.Subgroup, error) {
	rows, err := db.Query(`
SELECT id, title, day, time, group_id, subgroup_id
FROM lessons
WHERE group_id = ?
AND (subgroup_id IS NULL OR subgroup_id = ?)
`, groupID, userSubGroup_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subs []models.Subgroup

	for rows.Next() {
		var s models.Subgroup
		rows.Scan(&s.ID, &s.Name)
		subs = append(subs, s)
	}

	return subs, nil
}
