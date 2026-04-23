package logic

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
	"stud/models"
	"time"
)

var Lessons []models.Lesson

var lessID int = 1
var groupid int = 1

func AddLesson(l models.Lesson) error {

	for _, r := range Lessons {
		if r.Day == l.Day && r.Time == l.Time {
			return errors.New("err")
		}

	}

	l.ID = lessID
	lessID++
	l.GroupId = groupid
	Lessons = append(Lessons, l)

	return nil
}

func DeleteLesson(id int) {

	for i, l := range Lessons {
		if l.ID == id {
			Lessons = append(Lessons[:i], Lessons[i+1:]...)
			return
		}
	}

}

func ShowLessons() []models.Lesson {
	return Lessons
}

func TodayLessons() []models.Lesson {

	today := time.Now().Weekday().String()

	result := []models.Lesson{}

	for _, l := range Lessons {
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

	result := []models.Lesson{}

	for _, d := range Lessons {
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

	for i, l := range Lessons {

		if l.ID == id {

			switch field {

			case "1":
				Lessons[i].Title = value

			case "2":
				Lessons[i].Time = value

			case "3":
				Lessons[i].Day = value

			case "4":
				Lessons[i].Description = value

			default:
				return fmt.Errorf("invalid index")

			}
			return nil
		}
	}

	return fmt.Errorf("lesson not found")

}

func FilterListDay(day string) []models.Lesson {

	var result []models.Lesson

	for _, d := range Lessons {
		if strings.ToLower(d.Day) == strings.ToLower(day) {
			result = append(result, d)

		}
	}

	return result
}

func FilterListLess(less string) []models.Lesson {
	var result []models.Lesson
	for _, l := range Lessons {
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
	now := time.Now()

	var next *models.Lesson

	for _, l := range Lessons {

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
func SortTodayLesseonsTime() []models.Lesson {

	lessons := TodayLessons()

	sort.Slice(lessons, func(i, j int) bool {
		t1, _ := time.Parse("15:04", lessons[i].Time)
		t2, _ := time.Parse("15:04", lessons[j].Time)
		return t1.Before(t2)

	})
	return lessons
}

func UpdateLesson(updated models.Lesson) error {

	for i, l := range Lessons {

		if l.ID == updated.ID {

			if updated.Title != "" {
				Lessons[i].Title = updated.Title
			}

			if updated.Day != "" {
				Lessons[i].Day = updated.Day
			}

			if updated.Time != "" {
				Lessons[i].Time = updated.Time
			}

			if updated.Description != "" {
				Lessons[i].Description = updated.Description
			}

			return nil
		}
	}

	return errors.New("lesson not found")
}

//===================================================================================

func SaveToFile(lessons []models.Lesson, filepath string) error {
	file, err := os.Create(filepath)

	if err != nil {
		return err
	}
	defer file.Close()

	json.NewEncoder(file).Encode(lessons)
	return fmt.Errorf("lesson not found")
}

func LoadFromFile(filepath string) ([]models.Lesson, error) {

	file, err := os.Open(filepath)

	if err != nil {
		return []models.Lesson{}, err
	}
	defer file.Close()

	var lessons []models.Lesson

	json.NewDecoder(file).Decode(&lessons)
	return lessons, err
}
