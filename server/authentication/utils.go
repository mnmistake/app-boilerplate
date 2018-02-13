package authentication

import (
	"time"
	"errors"

	"golang.org/x/crypto/bcrypt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/raunofreiberg/kyrene/server"
	"github.com/raunofreiberg/kyrene/server/model"
)

func generateJWT(user interface{}) (string, error) {
	expireToken := time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &model.User{
		ID:       user.(model.User).ID,
		Username: user.(model.User).Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireToken,
		},
	})
	signedToken, err := token.SignedString(server.JwtSecret)

	if err != nil {
		return "", errors.New("Failed to generate token")
	}

	return signedToken, nil
}

func HashPassword(password string) ([]uint8, error) {
	bytePassword := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	return hashedPassword, nil
}
