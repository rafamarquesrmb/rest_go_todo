package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func ServerRun(port string, router *mux.Router) {
	fmt.Printf("Server is running on Port: %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}
