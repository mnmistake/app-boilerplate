package authentication

import (
	"errors"

	"github.com/raunofreiberg/kyrene/server/model"
)

func LoginUser(username string, password string) (interface{}, error) {
	queriedUser, err := QueryUser(username)
	if err != nil {
		return nil, errors.New("User not found")
	}

	isAuthenticated, err := IsAuthenticated(username, []byte(password))
	if err != nil {
		return nil, err
	}

	if isAuthenticated {
		signedToken, err := GenerateJWT(queriedUser)

		if err != nil {
			return nil, err
		}

		return model.Token{
			Token: signedToken,
		}, nil
	}

	return nil, nil
}

func RegisterUser(username string, password string) (interface{}, error) {
	queriedUser, err := QueryUser(username)
	if queriedUser != nil {
		return nil, errors.New("User already exists")
	}

	user, err := CreateUser(username, password)
	if err != nil {
		return nil, err
	}

	signedToken, err := GenerateJWT(user)
	if err != nil {
		return nil, err
	}

	return model.Token{
		Token: signedToken,
	}, nil
}
