package authentication

import (
	"fmt"

	"github.com/raunofreiberg/kyrene/server"
	"github.com/raunofreiberg/kyrene/server/model"
)

var (
	id          int
	username	string
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
		ID:          id,
		Username:    username,	
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
			ID:          id,
			Username:    username,	
		}, nil
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return nil, fmt.Errorf("User not found")
}