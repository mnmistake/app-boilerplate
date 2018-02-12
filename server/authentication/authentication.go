package authentication

import (
	"encoding/json"
	"fmt"
	"time"
	"net/http"

	"github.com/raunofreiberg/kyrene/server/model"
	"github.com/raunofreiberg/kyrene/server"
	jwt "github.com/dgrijalva/jwt-go"
)

func LoginFunc(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	fmt.Println(decoder)
}

func RegisterFunc(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	userData := struct {
		Username	string	`json:"username"`
		Password	string	`json:"password"`
	}{}
	err := decoder.Decode(&userData)

	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// if the user exists, don't create another one
	queriedUser, err := QueryUser(userData.Username)
	if queriedUser != nil {
		http.Error(w, "User already exists", http.StatusInternalServerError)
		return
	}

	user, err := CreateUser(userData.Username, userData.Password)	

	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	expireToken := time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &model.User{
		ID: user.(model.User).ID,
		Username: user.(model.User).Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt:	expireToken,
		},
	})
	signedToken, err := token.SignedString(server.JwtSecret)

	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}
	
	json.NewEncoder(w).Encode(struct{
		Token string `json:"token"`
	}{signedToken})
	return
}