package main

import (
	"log"

	"github.com/ptsgr/GoToDo"
	"github.com/ptsgr/GoToDo/pkg/handler"
)

func main() {
	srv := new(GoToDo.Server)
	handler := new(handler.Handler)

	if err := srv.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatalf("Error running http server: %s", err.Error())
	}
}
