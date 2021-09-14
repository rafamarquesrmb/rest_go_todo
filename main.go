package main

import (
	"fmt"

	app "github.com/rafamarquesrmb/rest_go_todo/src"
)

func main() {
	fmt.Println("App started ...")
	app.Init()
	defer fmt.Printf("App terminated...")
}
