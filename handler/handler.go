package handler

import (
	"encoding/json"
	"fmt"

	"net/http"
	"stud/models"
	"stud/service"
)

// http://localhost:8080/StudList

var lesson []models.Lesson

var fileName string

func Serv(FileName string) {

	fileName = FileName

	http.HandleFunc("/StudList/today", getToday)

	http.HandleFunc("/StudList", getLessons)

	http.HandleFunc("/StudList/add", addLesson)

	http.HandleFunc("/StudList/del", delLesson)

	http.HandleFunc("/StudList/update", updateLesson)

	http.HandleFunc("/StudList/day", DayLess)

	http.HandleFunc("/StudList/next", NextLess)

	http.HandleFunc("/StudList/nextToday", NextToday)

	http.ListenAndServe(":8080", nil)

	service.SaveToFile("data/" + fileName + ".json")

}

func getToday(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(service.TodayLessons())
}

func getLessons(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(service.ShowLessons())
}

func addLesson(w http.ResponseWriter, r *http.Request) {

	var newLesson models.Lesson

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&newLesson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	service.AddLesson(newLesson)
	service.SaveToFile("data/" + fileName + ".json")

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Lesson added"))
	fmt.Println(newLesson)

}

func delLesson(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	type deleteRequest struct {
		ID int
	}

	var req deleteRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	service.DeleteLesson(req.ID)
	service.SaveToFile("data/" + fileName + ".json")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Lesson deleted"))

}

func updateLesson(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var updated models.Lesson

	err := json.NewDecoder(r.Body).Decode(&updated)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err = service.UpdateLesson(updated)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	service.SaveToFile("data/" + fileName + ".json")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Lesson updated"))
}

func DayLess(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	day := r.URL.Query().Get("day")

	if day == "" {
		http.Error(w, "day is required", http.StatusBadRequest)
		return
	}

	result := service.DayLessons(day)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)

}

func NextLess(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	result := service.NextLesson()

	if result == nil {
		http.Error(w, "no next lesson", http.StatusNotFound)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)

}

func NextToday(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	result := service.NextLessonToday()

	if result == nil {
		http.Error(w, "no lessons today", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
