package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Base
	Username          string              `json:"username" gorm:"type:varchar(100);unique_index"`
	Password          string              `json:"-" gorm:"type:varchar(100)"`
	Group             []Group             `json:"-" gorm:"foreignkey:OwnerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Members           []Member            `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TranactionMembers []TransactionMember `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DebtsCreditor     []Debt              `json:"-" gorm:"foreignkey:CreditorID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DebtsDebtor       []Debt              `json:"-" gorm:"foreignkey:DebtorID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
