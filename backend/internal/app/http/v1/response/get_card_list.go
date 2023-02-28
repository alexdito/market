package response

import (
	"application/internal/app/src/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CardItemResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Code        string    `json:"code"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Sort        int       `json:"sort"`
}

type CardValueListStruct struct {
	C     *gin.Context
	Cards []entity.Card
}

type CardValueItemStruct struct {
	C *gin.Context
	entity.Card
}

func (s *CardValueListStruct) Response() []CardItemResponse {
	var response []CardItemResponse

	for _, card := range s.Cards {
		serializer := CardValueItemStruct{s.C, card}
		response = append(response, serializer.Response())
	}

	return response
}

func (s *CardValueItemStruct) Response() CardItemResponse {
	return CardItemResponse{
		ID:          s.ID,
		Name:        s.Name,
		Code:        s.Code,
		Description: s.Description,
		Image:       s.Image,
		Sort:        s.Sort,
	}
}
