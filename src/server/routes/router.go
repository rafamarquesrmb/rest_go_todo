package routes

import (
	"github.com/gorilla/mux"
)

func ConfigRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", GetAll).Methods("GET")
	router.HandleFunc("/{id}", GetById).Methods("GET")
	router.HandleFunc("/", Create).Methods("POST")
	router.HandleFunc("/{id}", Update).Methods("PUT")
	router.HandleFunc("/{id}", Delete).Methods("DELETE")
	return router
}
