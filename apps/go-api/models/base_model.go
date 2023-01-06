package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;" json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	base.ID = uuid.NewV4()

	return
}
