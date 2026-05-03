package repository

func GetOrCreateGroup(name string, course int, universityID int) (int, error) {
	var id int

	err := db.QueryRow(
		"SELECT id FROM groups WHERE name = ? AND course = ? AND university_id = ?",
		name, course, universityID,
	).Scan(&id)

	if err == nil {
		return id, nil
	}

	result, err := db.Exec(
		"INSERT INTO groups (name, course, university_id) VALUES (?, ?, ?)",
		name, course, universityID,
	)
	if err != nil {
		return 0, err
	}

	newID, _ := result.LastInsertId()
	return int(newID), nil
}
