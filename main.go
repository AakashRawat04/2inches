package main

import (
	"log"

	"github.com/aakashrawat04/2inches/server"
	"github.com/aakashrawat04/2inches/store"
)

func main() {
	store, err := store.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}
	server := server.NewAPIServer(":8080", store)
	server.Run()
}
