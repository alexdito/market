package entity

import (
	"github.com/google/uuid"
	"time"
)

func (model *Card) TableName() string {
	return "cards"
}

type Card struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Code        string    `db:"code"`
	Description string    `db:"description"`
	Image       string    `db:"image"`
	Sort        int       `db:"sort"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
