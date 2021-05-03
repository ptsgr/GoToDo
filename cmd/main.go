package main

import (
	"log"

	"github.com/ptsgr/GoToDo"
)

func main() {
	srv := new(GoToDo.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("Error running http server: %s", err.Error())
	}
}
