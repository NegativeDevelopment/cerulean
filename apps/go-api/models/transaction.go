package models

import (
	"github.com/google/uuid"
)

type Transaction struct {
	Base
	GroupID         uuid.UUID           `json:"groupId"`
	Amount          float64             `json:"amount" gorm:"type:decimal(10,2);not null;default:0.00;"`
	Currency        string              `json:"currency" gorm:"type:char(3);not null;default:'EUR';"`
	Title           string              `json:"titel" gorm:"type:varchar(255);not null;default:'';"`
	Description     string              `json:"description" gorm:"type:text;not null;default:'';"`
	CreatedByUserID uuid.UUID           `json:"-" gorm:"unique_index"`
	CreatedBy       User                `json:"createdBy" gorm:"foreignkey:CreatedByUserID;references:ID;"`
	Members         []TransactionMember `json:"members" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
