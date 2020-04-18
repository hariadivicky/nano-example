package main

import (
	"log"
	"todo/models"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	models.OpenDatabase("file:database/todo.db?cache=private")

	app := newApp()

	log.Println("server running at port 8080")

	err := app.Run(":8080")
	if err != nil {
		log.Fatalf("could not start application: %v", err)
	}
}
