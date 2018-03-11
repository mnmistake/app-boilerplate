package authentication

import (
	"errors"

	"github.com/raunofreiberg/kyrene/server/api/users"
	"github.com/raunofreiberg/kyrene/server/model"
)

func LoginUser(username string, password string) (interface{}, error) {
	queriedUser, err := users.QueryUser(username)
	if err != nil {
		return nil, errors.New("User not found")
	}

	isAuthenticated, err := users.IsAuthenticated(username, []byte(password))
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
	queriedUser, err := users.QueryUser(username)
	if queriedUser != nil {
		return nil, errors.New("User already exists")
	}

	hashedPassword, err := HashPassword(password)

	if err != nil {
		return nil, err
	}

	user, err := users.CreateUser(username, hashedPassword)
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
