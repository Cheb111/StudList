package models

type User struct {
	Id       int
	Name     string
	Password string
	UserRole string
	GroupId  int
	Subgroup_id int
	Gender   string
}
