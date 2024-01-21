package server

import (
	"log"

	"github.com/aakashrawat04/2inches/server/router"
	"github.com/aakashrawat04/2inches/store"
)

type APIServer struct {
	listenAddr string
}

func NewAPIServer(listenAddr string, database *store.PostgresStore) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Run() {
	r := router.RouterEngine()

	log.Println("Starting server on", s.listenAddr)
	r.Run(s.listenAddr)
}
