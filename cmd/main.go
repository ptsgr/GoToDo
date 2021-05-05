package main

import (
	"log"

	"github.com/ptsgr/GoToDo"
	"github.com/ptsgr/GoToDo/pkg/handler"
	"github.com/ptsgr/GoToDo/pkg/repository"
	"github.com/ptsgr/GoToDo/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handler := handler.NewHandler(services)
	srv := new(GoToDo.Server)

	if err := srv.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatalf("Error running http server: %s", err.Error())
	}
}
