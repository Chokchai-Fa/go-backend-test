package main

import (
	"backend_test/api"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	server, err := api.NewServer()
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start("0.0.0.0:8080")
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
