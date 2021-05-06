package main

import (
	"log"

	"github.com/ptsgr/GoToDo"
	"github.com/ptsgr/GoToDo/pkg/handler"
	"github.com/ptsgr/GoToDo/pkg/repository"
	"github.com/ptsgr/GoToDo/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error cannot initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handler := handler.NewHandler(services)
	srv := new(GoToDo.Server)

	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatalf("Error running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
