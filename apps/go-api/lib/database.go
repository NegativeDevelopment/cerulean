package lib

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDatabase() {
	var err error
	conn, err := gorm.Open(postgres.Open(os.Getenv("DB_STRING")), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB = conn
}
