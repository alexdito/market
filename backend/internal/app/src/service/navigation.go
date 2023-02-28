package service

import (
	"application/internal/app/common"
	"application/internal/app/repository"
	"application/internal/app/src/entity"
	"context"
	"net/http"
	"sync"
)

type NavigationService interface {
	GetValueList(ctx context.Context) ([]entity.Navigation, common.HttpError)
}

type navigationServiceImpl struct {
	mu                   *sync.RWMutex
	navigationRepository repository.NavigationRepository
}

func (service navigationServiceImpl) GetValueList(ctx context.Context) ([]entity.Navigation, common.HttpError) {
	list, err := service.navigationRepository.GetNavigationList(ctx)
	if err != nil {
		return list, common.NewHttpError(http.StatusInternalServerError, err)
	}
	return list, nil
}

func NewNavigationService(mu *sync.RWMutex, navigationRepository repository.NavigationRepository) NavigationService {
	return &navigationServiceImpl{
		mu:                   mu,
		navigationRepository: navigationRepository,
	}
}
