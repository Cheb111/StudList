package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"stud/models"
	"stud/repository"
	"stud/service"
)

var weeks = []string{
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
	"Sunday",
}

func main() {

	os.MkdirAll("data", os.ModePerm)

	scanner := bufio.NewScanner(os.Stdin)

	repository.InitDB()

	fmt.Println("Введите своё имя")

	scanner.Scan()

	name := scanner.Text()

	if name == "" {
		fmt.Println("Пустая строка")
		return
	}

	var gender string

	for {
		fmt.Println("Выберите пол")
		fmt.Println("1 - male")
		fmt.Println("2 - female")

		scanner.Scan()
		input := scanner.Text()
		choice, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Введите число 1 или 2:")
			continue
		}

		if choice == 1 {
			gender = "male"
			break
		} else if choice == 2 {
			gender = "female"
			break
		} else {
			fmt.Println("Неверный ввод:")
		}
	}

	var subgroups_id = 1
	
	fmt.Println("Введите пороль: ")

	scanner.Scan()

	password := scanner.Text()

	var userRole string

	if password == "" {
		userRole = "Guest"
	} else {
		userRole = "Admin"
	}

	fmt.Println("Введите название университета (полностью): ")
	scanner.Scan()
	UniName := scanner.Text()

	uniID, err := service.GetOrCreateUni(UniName)

	if err != nil {
		fmt.Println("error", err)
		return
	}

	fmt.Println("Введите название группы: ")
	scanner.Scan()
	group_name := scanner.Text()

	fmt.Println("Введите курс: ")

	scanner.Scan()
	courseStr := scanner.Text()
	course, err := strconv.Atoi(courseStr)

	groupId, err := service.GetOrCreateGroups(group_name, course, uniID)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	userID, err := service.RegisterUser(name, gender, subgroups_id, password, userRole, groupId, uniID)

	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println("Ваш Id", userID)

	//handler.Serv(FileName)

	for {
		//=========================================

		fmt.Println("")
		fmt.Println("1 - добавить")
		fmt.Println("2 - удалить")
		fmt.Println("3 - показать список")
		fmt.Println("4 - сегодня")
		fmt.Println("5 - редакт")
		fmt.Println("6 - выход")
		fmt.Println("")

		//=========================================

		fmt.Println("")
		fmt.Println("Введите что хотите сделать")

		scanner.Scan()
		todo := scanner.Text()

		todo = strings.ToLower(todo)

		switch todo {

		case "добавить", "1":

			if userRole == "Admin" {

				fmt.Println("")
				fmt.Println("Введите название предмета: ")

				scanner.Scan()
				title := scanner.Text()

				if err != nil {
					fmt.Println("error", err)
					return
				}

				fmt.Println("")
				fmt.Println("Выберите дату: ")

				for i, d := range weeks {
					fmt.Printf("%d - %s\n", i+1, d)
				}

				scanner.Scan()
				date := scanner.Text()

				index, _ := strconv.Atoi(date)

				if index < 1 || len(weeks) < index {
					fmt.Println("Невереный выбор!!!")
					continue
				}

				week := weeks[index-1]

				fmt.Println("")
				fmt.Println("Введите время в формате (xx:xx) : ")

				scanner.Scan()
				times := scanner.Text()

				_, err = time.Parse("15:04", times)

				if err != nil {
					fmt.Println("Неверный формат времени(xx:xx)")
					continue

				}

				fmt.Println("")
				fmt.Println("Введите описание(если надо): ")

				scanner.Scan()
				desc := scanner.Text()

				fmt.Println("")
				lesson := models.Lesson{
					User:        name,
					Title:       title,
					Day:         week,
					Time:        times,
					Description: desc,
					GroupId:     groupId,
				}

				err = service.AddLesson(lesson)

				if err != nil {
					fmt.Println(err)
					continue

				}
			} else {
				fmt.Println("У вас недостаточно прав(((")

				continue
			}

			continue

		//=================================================

		case "удалить", "2":

			if userRole == "Admin" {

				fmt.Println("Номер предмета")

				scanner.Scan()

				strid := scanner.Text()
				id, err := strconv.Atoi(strid)

				if err != nil {
					fmt.Println("Invalid type id")
					continue
				}

				service.DeleteLesson(id)
			} else {
				fmt.Println("У вас недостаточно прав(((")
				continue
			}

		//=================================================

		case "показать", "3":

			lessons := service.ShowLessons()

			if len(lessons) == 0 {
				fmt.Println("Список пуст")
				continue
			}

			service.SortTodayLessonsTime(lessons)

			fmt.Println("📚 Расписание:")

			for i, l := range lessons {
				fmt.Printf("%d. %s | %s | %s\n", i+1, l.Title, l.Day, l.Time)

				if l.Description != "" {
					fmt.Println("   📝", l.Description)
				}
			}

			fmt.Println("")
			fmt.Println("Введите что хотите сделать: ")

			fmt.Println("")
			fmt.Println("1 - Показать предметы на выбранный день")
			fmt.Println("2 - Показать когда предмет по дням")
			fmt.Println("3 - Закрыть список")
			fmt.Println("")

			scanner.Scan()
			id := scanner.Text()

			index, err := strconv.Atoi(id)

			if err != nil {
				fmt.Println("Введите число")
				continue
			}

			switch index {

			case 1:
				fmt.Println("")
				fmt.Println("Введите день:")
				fmt.Println("")
				scanner.Scan()
				day := scanner.Text()
				lessons := service.DayLessons(day)

				for _, d := range lessons {
					fmt.Println("")
					fmt.Printf("%d. %s | %s | %s | %s\n", d.ID, d.Title, d.Day, d.Time, d.Description)
					fmt.Println("")
				}

			case 2:
				fmt.Println("")
				fmt.Println("Введите предмет:")
				fmt.Println("")
				scanner.Scan()
				less := scanner.Text()
				lessons := service.FilterListLess(less)

				for _, l := range lessons {
					fmt.Println("")
					fmt.Printf("%d. %s | %s | %s | %s\n", l.ID, l.Title, l.Day, l.Time, l.Description)
					fmt.Println("")
				}

			case 3:
				continue

			}

		//=================================================

		case "сегодня", "4":

			lessons := service.TodayLessons()

			service.SortTodayLessonsTime(lessons)

			for _, l := range lessons {

				status := service.GetLessonStatus(l)

				var icon string

				switch status {
				case "first":
					icon = "🟣"
				case "past":
					icon = "🔴"
				case "current":
					icon = "🟢"
				case "future":
					icon = "🟡"
				}

				fmt.Printf("%s %d. %s | %s | %s\n", icon, l.ID, l.Title, l.Day, l.Time)
			}

			//=================================================

		case "редакт", "5":

			if userRole == "Admin" {
				fmt.Println("")
				fmt.Println("Введите номер задачи: ")

				scanner.Scan()
				id := scanner.Text()
				index, err := strconv.Atoi(id)

				if err != nil {
					fmt.Println("Ошибка")
					break
				}

				fmt.Println("")
				fmt.Println("\nВведите что хотите изменить: ")
				fmt.Println("")
				fmt.Println("1 - название")
				fmt.Println("2 - время")
				fmt.Println("3 - дата")
				fmt.Println("4 - описание")

				scanner.Scan()
				choice := scanner.Text()
				choice = strings.ToLower(choice)

				fmt.Println("")
				fmt.Println("Введите новое значени:")

				scanner.Scan()
				value := scanner.Text()

				service.EditLesson(index, choice, value)
			} else {
				fmt.Println("У вас недостаточно прав(((")
				continue
			}

		case "выход", "6":
			fmt.Println("")
			fmt.Println("Сохранить файл?")
			scanner.Scan()
			input := scanner.Text()
			input = strings.ToLower(input)
			if input == "нет" {

				//os.Remove("data/" + FileName + ".db")
				fmt.Println("Файл удалён!!!")
				fmt.Println("")
				fmt.Println("Пока!!!")
				return

			} else {
				fmt.Println("Файл сохранён!!!")
				fmt.Println("")
				fmt.Println("Пока!!!")

				return
			}

		}
	}

}
