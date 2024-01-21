package router

import (
	"github.com/aakashrawat04/2inches/store"
	"github.com/gin-gonic/gin"
)

type LinkService struct {
	store *store.PostgresStore
}

func (s *LinkService) getLinks(c *gin.Context) {
	links, err := s.store.GetLinks()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "error getting links",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "get links",
		"links":   links,
	})
}

func getLinkById(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "get link by id",
	})
}

func createLink(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "create link",
	})
}

func updateLink(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "update link",
	})
}

func deleteLink(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "delete link",
	})
}
