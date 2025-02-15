package main

import (
	"log"

	"github.com/Nurt0re/todo-app"
	"github.com/Nurt0re/todo-app/pkg/handler"
	"github.com/Nurt0re/todo-app/pkg/repository"
	"github.com/Nurt0re/todo-app/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running the server: %s", err.Error())
	}

}
