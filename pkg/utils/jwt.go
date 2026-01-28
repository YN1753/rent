package utils

import (
	"rent/internal/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(user model.UserInfo, secret string) (string, error) {
	key := []byte(secret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString(key)
}

func ParseToken(token string, secret string) (Claims, error) {
	var claims Claims
	key := []byte(secret)
	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return claims, err
	}
	return claims, nil
}
