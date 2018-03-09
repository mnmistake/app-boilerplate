package authentication

import (
	"errors"

	"github.com/raunofreiberg/kyrene/server/database"
	"github.com/raunofreiberg/kyrene/server/model"
	"golang.org/x/crypto/bcrypt"
)

var (
	id             int
	username       string
	hashedPassword []byte
)

func CreateUser(username string, password string) (interface{}, error) {
	hashedPassword, err := HashPassword(password)

	if err != nil {
		return nil, err
	}

	user := database.User{
		Username: username,
		Password: hashedPassword,
	}

	res := database.DB.Create(&user)

	if res.Error != nil {
		return nil, res.Error
	}

	return model.User{
		ID:       user.ID,
		Username: user.Username,
	}, nil
}

func QueryUser(username string) (interface{}, error) {
	user := database.User{}

	res := database.DB.Where("username = ?", username).First(&user)

	if res.Error != nil {
		return nil, res.Error
	}

	return model.User{
		ID:       user.ID,
		Username: user.Username,
	}, nil
}

func QueryUsers() (interface{}, error) {
	var users []model.User
	var dbUsers []database.User

	res := database.DB.Select("id, username").Find(&dbUsers)

	if res.Error != nil {
		return nil, res.Error
	}

	for _, x := range dbUsers {
		users = append(users, model.User{
			ID:       x.ID,
			Username: x.Username,
		})
	}

	return users, nil
}

func IsAuthenticated(username string, password []byte) (bool, error) {
	user := database.User{}

	res := database.DB.Where("username = ?", username).First(&user)

	if res.Error != nil {
		return false, res.Error
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, password); err != nil {
		return false, errors.New("Incorrect password")
	}

	return true, nil
}
