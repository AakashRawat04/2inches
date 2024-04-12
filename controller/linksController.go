package controller

import "github.com/gin-gonic/gin"

type LinksController struct {
}

func (l *LinksController) GetLinks(c *gin.Context) {
	c.JSON(200, gin.H{
		"links": "sooner or later, you will get the links here!",
	})
}
