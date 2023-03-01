package response

import (
	"application/internal/app/src/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type NavigationValueListStruct struct {
	C           *gin.Context
	Navigations []entity.Navigation
}

type NavigationValueItemStruct struct {
	C *gin.Context
	entity.Navigation
}

type NavigationValueItemResponse struct {
	ID uuid.UUID `json:"id"`
}

func (s *NavigationValueListStruct) Response() []NavigationValueItemResponse {
	var response []NavigationValueItemResponse
	for _, navigation := range s.Navigations {
		serializer := NavigationValueItemStruct{s.C, navigation}
		response = append(response, serializer.Response())
	}

	return response
}

func (s *NavigationValueItemStruct) Response() NavigationValueItemResponse {

	response := NavigationValueItemResponse{
		ID: s.ID,
	}
	return response
}
