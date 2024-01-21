package main

import (
	"github.com/aakashrawat04/2inches/server"
)

func main() {
	server := server.NewAPIServer(":8080")
	server.Run()
}
