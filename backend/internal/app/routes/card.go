package routes

import (
	"application/internal/app/http/v1/controller"
	"github.com/gin-gonic/gin"
)

func Card(router *gin.RouterGroup, controller controller.CardController) {
	publicV1 := router.Group("/v1/card")

	publicV1.GET("/list", controller.GetCardList)
}
