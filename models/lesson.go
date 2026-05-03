package models

import "database/sql"

type Lesson struct{
	ID int
	GroupId int
	Subgroup_id sql.NullInt64
	User string
	UserId int

	//WeekNum int
	
	Day string
	Time string
	Title string
	Description string

}