package server

import (
	"fmt"
	"net/http"
	"regexp"

	jwt "github.com/dgrijalva/jwt-go"
)

var JwtSecret = []byte("secret")

func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		bearerRegex, _ := regexp.Compile("(?:Bearer *)([^ ]+)(?: *)")
		bearerRegexMatches := bearerRegex.FindStringSubmatch(authorizationHeader)

		if len(bearerRegexMatches) == 0 {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		jwtToken := bearerRegexMatches[1]
		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return JwtSecret, nil
		})

		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		fmt.Println(token)

		next.ServeHTTP(w, r)
	})
}
