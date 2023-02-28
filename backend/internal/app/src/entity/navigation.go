package entity

import (
	"github.com/google/uuid"
	"time"
)

func (model *Navigation) TableName() string {
	return "navigation"
}

type Navigation struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Code        string    `db:"code"`
	Description string    `db:"description"`
	Sort        int       `db:"sort"`
	ParentUuid  uuid.UUID `db:"parent_uuid"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
