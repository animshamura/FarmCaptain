package services

import (
	"farmcaptain/models"
	"time"
	"github.com/dgrijalva/jwt-go"
	"os"
	"fmt"
)

// GenerateJWT generates a new JWT token for the farmer
func GenerateJWT(farmer models.Farmer) (string, error) {
	claims := jwt.MapClaims{
		"email": farmer.Email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
