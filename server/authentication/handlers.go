package authentication

import (
	"errors"

	"github.com/raunofreiberg/kyrene/server/database"
	"github.com/raunofreiberg/kyrene/server/model"
	"golang.org/x/crypto/bcrypt"
)

var db = database.Database()

func CreateUser(username string, password string) (interface{}, error) {
	hashedPassword, error := HashPassword(password)

	if error != nil {
		return nil, error
	}

	user := database.User{
		Username: username,
		Password: hashedPassword,
	}

	if _, err := db.Model(&user).Returning("id").Insert(); err != nil {
		return nil, err
	}

	return model.User{
		ID:       user.ID,
		Username: username,
	}, nil
}

func QueryUser(username string) (interface{}, error) {
	user := database.User{}

	_, err := db.QueryOne(
		&user,
		"SELECT id, username FROM users WHERE username = ?", username,
	)

	if err != nil {
		return nil, err
	}

	return model.User{
		ID:       user.ID,
		Username: user.Username,
	}, nil
}

func QueryUsers() (interface{}, error) {
	var users []model.User
	var dbUsers []database.User

	err := db.Model(&dbUsers).Select()

	if err != nil {
		return nil, err
	}

	for _, user := range dbUsers {
		users = append(users, model.User{
			ID:       user.ID,
			Username: user.Username,
		})
	}

	return users, nil
}

func IsAuthenticated(username string, password []byte) (bool, error) {
	user := database.User{}

	_, err := db.QueryOne(&user, "SELECT password FROM users WHERE username = ?", username)

	if err != nil {
		return false, nil
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, password); err != nil {
		return false, errors.New("Incorrect password")
	}

	return true, nil
}
