package authentication

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/raunofreiberg/kyrene/server"
	"github.com/raunofreiberg/kyrene/server/model"
	"golang.org/x/crypto/bcrypt"
)

// TODO: are tokens validated against their expiry date?
func GenerateJWT(user interface{}) (string, error) {
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
		return "", err
	}

	return signedToken, nil
}

func HashPassword(password string) ([]uint8, error) {
	bytePassword := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	return hashedPassword, nil
}
