package controller

import (
	"github.com/aakashrawat04/2inches/storage"
	"github.com/gin-gonic/gin"
)

type LinksController struct {
	storage *storage.LinkStorage
}

func NewLinksController(storage *storage.LinkStorage) *LinksController {
	return &LinksController{
		storage: storage,
	}
}

func (l *LinksController) GetLinks(c *gin.Context) {
	links, err := l.storage.GetLinks()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"links": links,
	})
}
