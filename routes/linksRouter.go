package routes

import (
	"github.com/aakashrawat04/2inches/controller"
	"github.com/gin-gonic/gin"
)

func LinkRoutes(app *gin.RouterGroup, linkController *controller.LinksController) {
	app.GET("/links", linkController.GetLinks)
}
