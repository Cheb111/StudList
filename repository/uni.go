package repository

import()

func GetOrCreateUni(name string)(int, error){
	var id int
	err := db.QueryRow(
		"SELECT id FROM universities WHERE name = ?",
		name,
	).Scan(&id)

	if err == nil{
		return id,nil
	}

	result, err := db.Exec(
		"INSERT INTO universities(name) VALUES (?)",
		name,
	)

	if err != nil{
		return 0,nil
	}

	newID, _ := result.LastInsertId()
	return int(newID), nil
}