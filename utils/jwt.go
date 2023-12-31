package utils

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var secret []byte

func GenerateToken(userId uint, email string) (string, error) {

	secretString := os.Getenv("JWT_SECRET")

	secret = []byte(secretString)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"email":  email,
	})

	return token.SignedString(secret)

}

func VerifyToken(token string) (uint, error) {

	secretString := os.Getenv("JWT_SECRET")

	secret = []byte(secretString)

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {

		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return secret, nil

	})

	if err != nil {
		return 0, errors.New("could not parse token")
	}

	if !parsedToken.Valid {
		return 0, errors.New("token is not valid")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("could not parse claims")
	}

	userId := claims["userId"].(float64)

	return uint(userId), nil

}
