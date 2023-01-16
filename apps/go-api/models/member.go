package models

import (
	"time"

	"github.com/google/uuid"
)

type Member struct {
	UserID    uuid.UUID  `json:"user_id" gorm:"primary_key;"`
	User      User       `json:"-" gorm:"foreignkey:UserID;references:ID;"`
	GroupID   uuid.UUID  `json:"group_id" gorm:"primary_key;"`
	Group     Group      `json:"-" gorm:"foreignkey:GroupID;references:ID;"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}
