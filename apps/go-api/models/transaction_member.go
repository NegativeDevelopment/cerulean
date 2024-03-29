package models

import (
	"time"

	"github.com/google/uuid"
)

type TransactionMember struct {
	TransactionID uuid.UUID   `json:"transactionId" gorm:"primary_key;"`
	Transaction   Transaction `json:"-" gorm:"foreignkey:TransactionID;references:ID;"`
	UserID        uuid.UUID   `json:"userId" gorm:"primary_key;"`
	User          User        `json:"-" gorm:"foreignkey:UserID;references:ID;"`
	CreatedAt     time.Time   `json:"-"`
	UpdatedAt     time.Time   `json:"-"`
	DeletedAt     *time.Time  `json:"-" sql:"index"`
}
