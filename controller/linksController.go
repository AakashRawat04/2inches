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

func (l *LinksController) GetLinksByUser(c *gin.Context) {
	userID := c.Param("userid")
	links, err := l.storage.GetAllLinksByUserId(userID)
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

func (l *LinksController) GetLinkByShortCode(c *gin.Context) {
	shortCode := c.Param("short_code")
	link, err := l.storage.GetLinkByShortCode(shortCode)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"link": link,
	})
}

func (l *LinksController) GetActiveLinksByUser(c *gin.Context) {
	userID := c.Param("userid")
	links, err := l.storage.GetActiveLinksByUserId(userID)
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

func (l *LinksController) GetDisabledLinksByUser(c *gin.Context) {
	userID := c.Param("userid")
	links, err := l.storage.GetDisabledLinksByUserId(userID)
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

func (l *LinksController) GetArchivedLinksByUser(c *gin.Context) {
	userID := c.Param("userid")
	links, err := l.storage.GetArchivedLinksByUserId(userID)
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

func (l *LinksController) DisableLink(c *gin.Context) {
	shortCode := c.Param("short_code")
	err := l.storage.DisableLink(shortCode)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{})
}
