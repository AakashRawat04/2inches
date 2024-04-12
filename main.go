package main

import (
	"github.com/aakashrawat04/2inches/controller"
	"github.com/aakashrawat04/2inches/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	api := r.Group("/api")
	routes.LinkRoutes(api, &controller.LinksController{})
	r.Run(":3000")
}
