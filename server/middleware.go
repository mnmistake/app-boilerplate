package server

import (
	"regexp"
	"context"
	"net/http"
)

func PassJwtContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		bearerRegex, _ := regexp.Compile("(?:Bearer *)([^ ]+)(?: *)")
		bearerRegexMatches := bearerRegex.FindStringSubmatch(authorizationHeader)
		ctx := context.WithValue(r.Context(), "jwt", "")

		if len(bearerRegexMatches) != 0 {
			jwtToken := bearerRegexMatches[1]
			ctx = context.WithValue(r.Context(), "jwt", jwtToken)
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
