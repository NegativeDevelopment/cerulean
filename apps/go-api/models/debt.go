package models

import (
	"time"

	"github.com/google/uuid"
)

type Debt struct {
	CreditorID uuid.UUID  `json:"creditor_id" gorm:"primary_key;"`
	Creditor   User       `json:"-"`
	DebtorID   uuid.UUID  `json:"deptor_id" gorm:"primary_key;"`
	Debtor     User       `json:"-"`
	GroupID    uuid.UUID  `json:"group_id" gorm:"primary_key;"`
	Group      Group      `json:"-" gorm:"foreignkey:GroupID;references:ID;"`
	Amount     float64    `json:"amount" gorm:"type:decimal(10,2);not null;default:0.00;"`
	Currency   string     `json:"currency" gorm:"type:char(3);not null;default:'EUR';"`
	CreatedAt  time.Time  `json:"-"`
	UpdatedAt  time.Time  `json:"-"`
	DeletedAt  *time.Time `json:"-" sql:"index"`
}
