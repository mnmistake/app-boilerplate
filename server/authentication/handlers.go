package authentication

import (
	"errors"

	"github.com/raunofreiberg/kyrene/server"
	"github.com/raunofreiberg/kyrene/server/model"
	"golang.org/x/crypto/bcrypt"
)

var (
	id             int
	username       string
	hashedPassword []byte
)

func CreateUser(username string, password string) (interface{}, error) {
	hashedPassword, error := HashPassword(password)
	if error != nil {
		panic(error)
	}

	err := server.DB.QueryRow(
		"INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id",
		username,
		hashedPassword,
	).Scan(&id)

	if err != nil {
		panic(err)
	}

	return model.User{
		ID:       id,
		Username: username,
	}, nil
}

func QueryUser(username string) (interface{}, error) {
	rows, err := server.DB.Query("SELECT id, username FROM users WHERE username=$1", username)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		err := rows.Scan(&id, &username)

		if err != nil {
			panic(err)
		}

		return model.User{
			ID:       id,
			Username: username,
		}, nil
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return nil, errors.New("User not found")
}

func AuthenticateUser(username string, password []byte) (bool, error) {
	queryErr := server.DB.QueryRow(
		"SELECT password FROM users where username=$1",
		username,
	).Scan(&hashedPassword)

	if queryErr != nil {
		panic(queryErr)
	}

	err := bcrypt.CompareHashAndPassword(hashedPassword, password)

	if err != nil {
		return false, errors.New("Incorrect password")
	}

	return true, nil
}
