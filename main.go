package main

import (
	"log"

	"github.com/aakashrawat04/2inches/database"
)

func main() {
	store, err := database.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}
	server := NewAPIServer(":8080", store)
	server.Run()
}
