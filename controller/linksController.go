package controller

import (
	"github.com/aakashrawat04/2inches/models"
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

func (l *LinksController) CreateLink(c *gin.Context) {
	var request models.CreateLinkRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	shortCode, err := l.storage.CreateLink(&request)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"short_code": shortCode,
	})
}
