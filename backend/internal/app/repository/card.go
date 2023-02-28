package repository

import (
	"application/internal/app/src/entity"
	"context"
	"gorm.io/gorm"
)

type CardRepository interface {
	GetCardList(ctx context.Context) ([]entity.Card, error)
}

type cardRepositoryImpl struct {
	db *gorm.DB
}

func (repository cardRepositoryImpl) GetCardList(ctx context.Context) ([]entity.Card, error) {
	var cards []entity.Card

	result := repository.db.Find(&cards).Order("sort").Limit(12)

	if result.Error != nil {
		panic("Cards not found")
	}

	return cards, nil
}

func NewCardRepository(db *gorm.DB) CardRepository {
	return &cardRepositoryImpl{
		db: db,
	}
}
