package lib

import (
	"os"
	"strconv"
	"time"

	"github.com/NegativeDevelopment/cerulean/go-api/models"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(user *models.User) (string, error) {
	tokenLifespanHours, err := strconv.Atoi(os.Getenv("JWT_LIFESPAN_HOURS"))

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"exp":    time.Now().Add(time.Hour * time.Duration(tokenLifespanHours)).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
