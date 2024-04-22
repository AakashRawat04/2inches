package routes

import (
	"github.com/aakashrawat04/2inches/controller"
	"github.com/gin-gonic/gin"
)

func LinkRoutes(app *gin.RouterGroup, linkController *controller.LinksController) {
	app.GET("/links", linkController.GetLinks)
	app.POST("/tiny", linkController.CreateLink)
	app.GET("/getAllLinks/:userid", linkController.GetLinksByUser)
	app.GET("/link/:short_code", linkController.GetLinkByShortCode)
	app.GET("/getActiveLinks/:userid", linkController.GetActiveLinksByUser)
	app.GET("/getDisabledLinks/:userid", linkController.GetDisabledLinksByUser)
	app.GET("/getArchivedLinks/:userid", linkController.GetArchivedLinksByUser)
	app.PUT("/disableLink/:short_code", linkController.DisableLink)
}
