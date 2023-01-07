package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Base
	Name     string `gorm:"unique" binding:"required" json:"name"`
	Password string `binding:"required" json:"password"`
}

func (user *User) HashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hash)
	return nil
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
