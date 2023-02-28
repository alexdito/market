package service

import (
	"application/internal/app/common"
	"application/internal/app/repository"
	"application/internal/app/src/entity"
	"context"
	"net/http"
	"sync"
)

type CardService interface {
	GetValueList(ctx context.Context) ([]entity.Card, common.HttpError)
}

type cardServiceImpl struct {
	mu             *sync.RWMutex
	cardRepository repository.CardRepository
}

func (service cardServiceImpl) GetValueList(ctx context.Context) ([]entity.Card, common.HttpError) {
	list, err := service.cardRepository.GetCardList(ctx)
	if err != nil {
		return list, common.NewHttpError(http.StatusInternalServerError, err)
	}
	return list, nil
}

func NewCardService(mu *sync.RWMutex, CardRepository repository.CardRepository) CardService {
	return &cardServiceImpl{
		mu:             mu,
		cardRepository: CardRepository,
	}
}
