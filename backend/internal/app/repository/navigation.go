package repository

import (
	"application/internal/app/src/entity"
	"context"
	"gorm.io/gorm"
)

type NavigationRepository interface {
	GetNavigationList(ctx context.Context) ([]entity.Navigation, error)
}

type navigationRepositoryImpl struct {
	db *gorm.DB
}

func (repository navigationRepositoryImpl) GetNavigationList(ctx context.Context) ([]entity.Navigation, error) {
	var navigations []entity.Navigation

	result := repository.db.Find(&navigations).Order("sort")

	if result.Error != nil {
		panic("Navigation not found")
	}

	return navigations, nil
}

func NewNavigationRepository(db *gorm.DB) NavigationRepository {
	return &navigationRepositoryImpl{
		db: db,
	}
}
