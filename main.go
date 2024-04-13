package main

import (
	"github.com/aakashrawat04/2inches/controller"
	"github.com/aakashrawat04/2inches/db"
	"github.com/aakashrawat04/2inches/routes"
	"github.com/aakashrawat04/2inches/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	db := db.NewPostgresConnection()
	linkStorage := storage.NewLinkStorage(db)
	linkController := controller.NewLinksController(linkStorage)
	api := r.Group("/api")
	routes.LinkRoutes(api, linkController)
	r.Run(":3000")
}
