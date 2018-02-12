package authentication

import (
	"encoding/json"
	"fmt"
	"errors"
	"time"
	"net/http"

	"github.com/raunofreiberg/kyrene/server/model"
	"github.com/raunofreiberg/kyrene/server"
	jwt "github.com/dgrijalva/jwt-go"
)

func generateJWT(user interface{}) (string, error) {
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
		return "", errors.New("Failed to generate token")
	}

	return signedToken, nil
}

func LoginFunc(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	userData := model.UserData{}
	err := decoder.Decode(&userData)

	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	queriedUser, err := QueryUser(userData.Username)
	if err != nil {
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}

	isAuthenticated, err := LoginUser(userData.Username, []byte(userData.Password))
	if err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	if isAuthenticated {
		signedToken, err := generateJWT(queriedUser)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		json.NewEncoder(w).Encode(model.Token{
			Token: signedToken,
		})
		return
	}
}

func RegisterFunc(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	userData := model.UserData{}
	err := decoder.Decode(&userData)

	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

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

	signedToken, err := generateJWT(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	json.NewEncoder(w).Encode(model.Token{
		Token: signedToken,
	})
	return
}