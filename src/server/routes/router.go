package routes

import (
	"github.com/gorilla/mux"
	"github.com/rafamarquesrmb/rest_go_todo/src/controllers"
)

func ConfigRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", controllers.GetAll).Methods("GET")
	router.HandleFunc("/{id}", controllers.GetById).Methods("GET")
	router.HandleFunc("/", controllers.Create).Methods("POST")
	router.HandleFunc("/{id}", controllers.Update).Methods("PUT")
	router.HandleFunc("/{id}", controllers.Delete).Methods("DELETE")
	router.HandleFunc("/completed/", controllers.GetAllCompleted).Methods("GET")
	router.HandleFunc("/complete/{id}", controllers.Completer).Methods("GET")
	router.HandleFunc("/notcompleted/", controllers.GetAllNotCompleted).Methods("GET")
	return router
}
