package router

import (
	"github.com/aakashrawat04/2inches/store"
	"github.com/gin-gonic/gin"
)

func RouterEngine() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")

	store, err := store.NewPostgresStore()
	if err != nil {
		panic(err)
	}

	linkService := &LinkService{store: store}

	v1.GET("/links", linkService.getLinks)
	v1.GET("/links/:id", getLinkById)
	v1.POST("/links", createLink)
	v1.PUT("/links/:id", updateLink)
	v1.DELETE("/links/:id", deleteLink)

	return router
}
