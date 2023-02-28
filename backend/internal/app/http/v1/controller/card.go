package controller

import (
	"application/internal/app/http/v1/response"
	"application/internal/app/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CardController interface {
	GetCardList(ctx *gin.Context)
}

func NewCardController(service service.CardService) CardController {
	return &cardControllerImpl{
		service: service,
	}
}

type cardControllerImpl struct {
	service service.CardService
}

func (controller *cardControllerImpl) GetCardList(ctx *gin.Context) {
	cardList, err := controller.service.GetValueList(ctx)

	if err != nil {
		ctx.JSON(err.HttpCode(), gin.H{"message": err.Error()})
		return
	}

	serializer := response.CardValueListStruct{C: ctx, Cards: cardList}

	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	ctx.JSON(http.StatusOK, gin.H{"data": serializer.Response()})
}
