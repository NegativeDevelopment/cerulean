package models

type User struct {
	Base
	Name string `binding:"required" json:"name"`
}
