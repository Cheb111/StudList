package repository

import "stud/models"

var Lessons []models.Lesson

func GetLessons()[]models.Lesson{
	return Lessons
}

func SetLessons(l []models.Lesson){
	Lessons = l
}