package routes

import (
	"application/internal/app/http/v1/controller"
	"github.com/gin-gonic/gin"
)

func Navigation(router *gin.RouterGroup, controller controller.NavigationController) {
	publicV1 := router.Group("/v1/navigation")

	publicV1.GET("/list", controller.GetNavigationValueList)
	publicV1.GET("/main-structure", controller.GetMainNavigationStructure)
}
