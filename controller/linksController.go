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
	something := l.storage.GetLinks()
	if something != nil {
		c.JSON(500, gin.H{
			"error": "Lmao, something went wrong!",
		})
		return
	}
	c.JSON(200, gin.H{
		"links": "sooner or later, you will get the links here!",
	})
}
