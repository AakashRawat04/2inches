package main

import (
	"log"
	"net/http"

	"github.com/aakashrawat04/2inches/database"
	"github.com/gin-gonic/gin"
)

type APIServer struct {
	listenAddr string
	database   database.Database
}

func NewAPIServer(listenAddr string, database database.Database) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		database:   database,
	}
}

func (s *APIServer) Run() {
	r := gin.Default()
	r.GET("/links", s.getLinks)

	log.Println("Starting server on", s.listenAddr)
	r.Run(s.listenAddr)
}

func (s *APIServer) getLinks(c *gin.Context) {
	links, err := s.database.GetLinks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"links": links,
	})
}
