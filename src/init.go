package src

import (
	"github.com/rafamarquesrmb/rest_go_todo/src/database"
	"github.com/rafamarquesrmb/rest_go_todo/src/server"
	"github.com/rafamarquesrmb/rest_go_todo/src/server/routes"
)

func Init() {
	port := ":8000"
	router := routes.ConfigRoutes()
	server.ServerRun(port, router)
	database.StartDB()

}
