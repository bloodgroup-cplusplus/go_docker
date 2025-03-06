package main

import (
	"log"

	"github.com/bloodgroup-cplusplus/go_docker/cmd/api"
)

func main () {
	server := api.NewAPIServer(":8080",nil)

	if err := server.Run(); err !=nil {
		log.Fatalf(err)
	}
}
