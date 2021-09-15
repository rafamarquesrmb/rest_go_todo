package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/rafamarquesrmb/rest_go_todo/src/database"
	"github.com/rafamarquesrmb/rest_go_todo/src/models"
)

func GetAll(w http.ResponseWriter, _ *http.Request) {
	db := database.GetDatabase()
	w.Header().Set("Content-type", "application/json")
	tasks := []models.Task{}

	err := db.Find(&tasks).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	json.NewEncoder(w).Encode(tasks)

}
func GetById(w http.ResponseWriter, r *http.Request) {
	db := database.GetDatabase()
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	var task models.Task
	err := db.First(&task, id).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	json.NewEncoder(w).Encode(task)
}

func Create(w http.ResponseWriter, r *http.Request) {
	db := database.GetDatabase()
	w.Header().Set("Content-type", "application/json")
	var task models.Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	err := db.Create(&task).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := database.GetDatabase()
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	var task models.Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	if task.ID != 0 {
		if task.ID != id {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
	task.ID = id
	rows, _ := db.Model(&models.Task{}).Where("id = ?", id).Rows() // (*sql.Rows, error)
	var before_up_task models.Task
	defer rows.Close()
	for rows.Next() {
		db.ScanRows(rows, &before_up_task)
	}
	task.CreatedAt = before_up_task.CreatedAt
	if before_up_task.Completed {
		task.Completed = before_up_task.Completed
		task.CompletedAt = before_up_task.CompletedAt
	} else {
		if task.Completed {
			task.CompletedAt = time.Now()
		} else {
			task.Completed = before_up_task.Completed
			task.CompletedAt = before_up_task.CompletedAt
		}
	}
	err = db.Save(&task).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := database.GetDatabase()
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	var task models.Task
	err := db.Delete(&task, id).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	w.WriteHeader(http.StatusOK)
}

func Completer(w http.ResponseWriter, r *http.Request) {
	db := database.GetDatabase()
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	// id := params["id"]
	rows, _ := db.Model(&models.Task{}).Where("id = ?", params["id"]).Rows() // (*sql.Rows, error)
	var task models.Task
	defer rows.Close()
	for rows.Next() {
		// ScanRows scan a row into user
		db.ScanRows(rows, &task)

		// do something
	}
	// err := db.First(&task, id).Error
	task.Completed = true
	task.CompletedAt = time.Now()
	err := db.Save(&task).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

func GetAllCompleted(w http.ResponseWriter, _ *http.Request) {
	db := database.GetDatabase()
	w.Header().Set("Content-type", "application/json")
	tasks := []models.Task{}

	err := db.Find(&tasks, "completed = ?", true).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} // (*sql.Rows, error)

	json.NewEncoder(w).Encode(tasks)
}

func GetAllNotCompleted(w http.ResponseWriter, _ *http.Request) {
	db := database.GetDatabase()
	w.Header().Set("Content-type", "application/json")
	tasks := []models.Task{}

	err := db.Find(&tasks, "completed = ?", false).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} // (*sql.Rows, error)

	json.NewEncoder(w).Encode(tasks)
}
