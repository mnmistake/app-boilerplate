package server

import (
	"errors"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

func ValidateJWT(jwtToken string) (bool, error) {
	if len(jwtToken) == 0 {
		return false, errors.New("Missing token")
	}

	_, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return false, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return JwtSecret, nil
	})

	if err != nil {
		return false, errors.New("Invalid token")
	}

	return true, nil
}
