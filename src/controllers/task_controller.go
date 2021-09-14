package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rafamarquesrmb/rest_go_todo/src/database"
	"github.com/rafamarquesrmb/rest_go_todo/src/models"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	db := database.GetDatabase()
	tasks := []models.Task{}

	err := db.Find(&tasks).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(tasks)

}
func GetById(w http.ResponseWriter, r *http.Request) {
	db := database.GetDatabase()
	params := mux.Vars(r)
	id := params["id"]
	var task models.Task
	err := db.First(&task, id).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(task)
}

func Create(w http.ResponseWriter, r *http.Request) {
	db := database.GetDatabase()
	var task models.Task
	_ = json.NewDecoder(r.Body).Decode(task)
	err := db.Create(&task).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := database.GetDatabase()
	var task models.Task
	_ = json.NewDecoder(r.Body).Decode(task)
	err := db.Save(&task).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := database.GetDatabase()
	params := mux.Vars(r)
	id := params["id"]
	var task models.Task
	err := db.Delete(&task, id).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}
