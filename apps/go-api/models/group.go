package models

import "github.com/google/uuid"

type Group struct {
	Base
	Name    string    `json:"name" gorm:"type:varchar(100);required"`
	OwnerID uuid.UUID `json:"owner_id"`
	Members []Member  `json:"members" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (group *Group) IsMember(userID uuid.UUID) bool {
	for _, member := range group.Members {
		if member.UserID == userID {
			return true
		}
	}
	return false
}
