package main

import (
	"todo/controllers"

	"github.com/hariadivicky/nano"
)

// newApp returns nano Engine instance.
func newApp() *nano.Engine {
	n := nano.New()

	n.Use(nano.Recovery())

	// Allow all origins
	n.Use(nano.CORSWithConfig(nano.CORSConfig{
		AllowedOrigins: []string{"*"},
	}))

	todo := new(controllers.Todo)

	n.GET("/todos", todo.Index)
	n.POST("/todos", todo.Store)
	n.GET("/todos/:id", todo.Show)
	n.DELETE("/todos/:id", todo.Destroy)
	n.PUT("/todos/:id/toggle", todo.Toggle)

	return n
}
