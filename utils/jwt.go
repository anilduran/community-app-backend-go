package utils

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var secret []byte

func GenerateToken(userId uuid.UUID, email string) (string, error) {

	secretString := os.Getenv("JWT_SECRET")

	secret = []byte(secretString)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId.String(),
		"email":  email,
	})

	return token.SignedString(secret)

}

func VerifyToken(token string) (uuid.UUID, error) {

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
		return uuid.Nil, errors.New("could not parse token")
	}

	if !parsedToken.Valid {
		return uuid.Nil, errors.New("token is not valid")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return uuid.Nil, errors.New("could not parse claims")
	}

	userId, _ := uuid.Parse(claims["userId"].(string))

	return userId, nil

}
