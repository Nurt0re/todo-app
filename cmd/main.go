package main

import (
	"log"

	"github.com/Nurt0re/todo-app"
	"github.com/Nurt0re/todo-app/pkg/handler"
)

func main() {
	handler:= new(handler.Handler)

	srv := new(todo.Server)
	if err:=srv.Run("8080", handler.InitRoutes()); err!=nil{
		log.Fatalf("error occured while running the server: %s", err.Error())
	}

}
