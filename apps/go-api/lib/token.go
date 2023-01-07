package lib

import (
	"os"
	"strconv"
	"time"

	"github.com/NegativeDevelopment/cerulean/go-api/models"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWTToken(user *models.User) (string, error) {
	tokenExpireTime, err := strconv.Atoi(os.Getenv("TOKEN_EXPIRE_TIME_HOURS"))

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"exp":    time.Now().Add(time.Hour * time.Duration(tokenExpireTime)).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
