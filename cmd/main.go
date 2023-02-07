package main

import (
	"log"

	"github.com/Polonkoev/todo-app"
	"github.com/Polonkoev/todo-app/package/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
