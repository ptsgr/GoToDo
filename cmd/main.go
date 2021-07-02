package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/ptsgr/GoToDo"
	"github.com/ptsgr/GoToDo/pkg/handler"
	"github.com/ptsgr/GoToDo/pkg/repository"
	"github.com/ptsgr/GoToDo/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error cannot initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Errorf("error loading env variables: %s", err.Error())
	}

	if "" == os.Getenv("DB_PASSWORD") {
		logrus.Fatalf("error cannot initializing database password")
	}

	db, err := repository.NewPostgresDB(repository.DBConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("failed to connect Database: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := handler.NewHandler(services)
	srv := new(GoToDo.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
			logrus.Fatalf("error running http server: %s", err.Error())
		}

	}()

	logrus.Print("GoToDo Started")
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	close(quit)

	logrus.Print("GoToDo Shutting Down")
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occurred on server shutting down: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("error occurred on db connection close: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
