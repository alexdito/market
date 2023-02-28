package controller

import (
	"application/internal/app/http/v1/response"
	"application/internal/app/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type NavigationController interface {
	GetNavigationValueList(ctx *gin.Context)
	GetMainNavigationStructure(ctx *gin.Context)
}

func NewNavigationController(service service.NavigationService) NavigationController {
	return &navigationControllerImpl{
		service: service,
	}
}

type navigationControllerImpl struct {
	service service.NavigationService
}

func (controller *navigationControllerImpl) GetNavigationValueList(ctx *gin.Context) {
	directionsValueList, err := controller.service.GetValueList(ctx)

	if err != nil {
		ctx.JSON(err.HttpCode(), gin.H{"message": err.Error()})
		return
	}

	serializer := response.NavigationValueListStruct{C: ctx, Navigations: directionsValueList}

	ctx.JSON(http.StatusOK, gin.H{"data": serializer.Response()})
}

func (controller *navigationControllerImpl) GetMainNavigationStructure(ctx *gin.Context) {
	navigationStructureList, err := controller.service.GetValueList(ctx)

	if err != nil {
		ctx.JSON(err.HttpCode(), gin.H{"message": err.Error()})
		return
	}

	serializer := response.NavigationListStruct{C: ctx, Navigations: navigationStructureList}

	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	ctx.JSON(http.StatusOK, gin.H{"data": serializer.Response()})
}
