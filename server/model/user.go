package model

import (
	"github.com/dgrijalva/jwt-go"
)

type User struct {
	ID       int    `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"-"`
	jwt.StandardClaims
}

type UserData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
