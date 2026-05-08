package service

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"stud/models"
	"stud/repository"
	"time"
)

var lessID int = 1
var groupid int = 1

func AddLesson(l models.Lesson) error {

	return repository.AddLesson(l)

}


func DeleteLesson(id int) error {

	lessons := repository.GetLessons()

	for i, l := range lessons {
		if l.ID == id {
			lessons = append(lessons[:i], lessons[i+1:]...)

			return repository.DeleteLesson(id)
		}
	}
	return fmt.Errorf("error")

}

func ShowLessons() []models.Lesson {
	return repository.GetLessons()
}

func TodayLessons() []models.Lesson {

	lessons := repository.GetLessons()

	today := time.Now().Weekday().String()

	result := []models.Lesson{}

	for _, l := range lessons {
		if strings.ToLower(l.Day) == strings.ToLower(today) {
			result = append(result, l)
		}

	}

	if result == nil {
		return result

	}
	return result
}

// =============================================
func DayLessons(day string) []models.Lesson {

	lessons := repository.GetLessons()

	result := []models.Lesson{}

	for _, d := range lessons {
		if strings.ToLower(d.Day) == strings.ToLower(day) {
			result = append(result, d)
		}
	}
	if result == nil {
		return result
	}
	return result
}

//=============================================

func EditLesson(id int, field string, value string) error {

	lessons := repository.GetLessons()

	for i, l := range lessons {

		if l.ID == id {

			switch field {

			case "1":
				lessons[i].Title = value

			case "2":
				lessons[i].Time = value

			case "3":
				lessons[i].Day = value

			case "4":
				lessons[i].Description = value

			default:
				return fmt.Errorf("invalid index")

			}
			return nil
		}
	}

	return fmt.Errorf("lesson not found")

}

func FilterListLess(less string) []models.Lesson {
	lessons := repository.GetLessons()
	var result []models.Lesson
	for _, l := range lessons {
		if strings.ToLower(l.Title) == strings.ToLower(less) {
			result = append(result, l)

		}

	}
	return result

}

func TimeParse(t string) (time.Time, error) {
	return time.Parse("15:04", t)
}

func GetLessonStatus(lesson models.Lesson) string {
	now := time.Now()

	start, err := TimeParse(lesson.Time)

	if err != nil {
		return "unknown"
	}

	startTime := time.Date(
		now.Year(), now.Month(), now.Day(),
		start.Hour(), start.Minute(), 0, 0, now.Location(),
	)

	endTime := startTime.Add(90 * time.Minute)
	if now.After(endTime) {
		return "past"
	} else if now.After(startTime) && now.Before(endTime) {
		return "current"
	} else {
		return "future"
	}
}

func parseTime(t string, now time.Time) time.Time {
	parsed, _ := TimeParse(t)

	return time.Date(
		now.Year(), now.Month(), now.Day(),
		parsed.Hour(), parsed.Minute(), 0, 0, now.Location(),
	)
}

func NextLesson() *models.Lesson {
	lessons := repository.GetLessons()
	now := time.Now()

	var next *models.Lesson

	for _, l := range lessons {

		start, err := TimeParse(l.Time)
		if err != nil {
			continue
		}

		startTime := time.Date(
			now.Year(), now.Month(), now.Day(),
			start.Hour(), start.Minute(), 0, 0, now.Location(),
		)

		if startTime.After(now) {
			if next == nil || startTime.Before(parseTime(next.Time, now)) {
				tmp := l
				next = &tmp
			}
		}
	}

	return next
}

func NextLessonToday() *models.Lesson {
	now := time.Now()
	today := now.Weekday().String()

	lessons := DayLessons(today)

	var next *models.Lesson

	for _, l := range lessons {

		start, err := TimeParse(l.Time)
		if err != nil {
			continue
		}

		startTime := time.Date(
			now.Year(), now.Month(), now.Day(),
			start.Hour(), start.Minute(), 0, 0, now.Location(),
		)

		if startTime.After(now) {
			if next == nil || startTime.Before(parseTime(next.Time, now)) {
				tmp := l
				next = &tmp
			}

		}
	}

	return next
}

func SortTodayLessonsTime(lessons []models.Lesson) {

	sort.Slice(lessons, func(i, j int) bool {
		t1, _ := time.Parse("15:04", lessons[i].Time)
		t2, _ := time.Parse("15:04", lessons[j].Time)
		return t1.Before(t2)

	})
}

func UpdateLesson(updated models.Lesson) error {

	lessons := repository.GetLessons()

	for i, l := range lessons {

		if l.ID == updated.ID {

			if updated.Title != "" {
				lessons[i].Title = updated.Title
			}

			if updated.Day != "" {
				lessons[i].Day = updated.Day
			}

			if updated.Time != "" {
				lessons[i].Time = updated.Time
			}

			if updated.Description != "" {
				lessons[i].Description = updated.Description
			}

			return nil
		}
	}

	return errors.New("lesson not found")
}

//===================================================================================

//====================================================

func RegisterUser(name string, gender string, subgroups_id int, password string, userRole string, groupId int, UniID int) (int, error) {
	return repository.CreateUser(name, gender,subgroups_id , password, userRole, groupId, UniID)
}

func GetOrCreateGroups(name string, course int, uniId int)(int, error){
	return repository.GetOrCreateGroup(name, course, uniId)
}

func GetOrCreateUni(name string)(int, error){
	return repository.GetOrCreateUni(name)
}

//====================================================
