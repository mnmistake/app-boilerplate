package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-test/deep"
	"github.com/raunofreiberg/kyrene/server"
)

type ResponseMock struct {
	Name string
}

func TestJwtMiddleware1(t *testing.T) {
	req, err := http.NewRequest("GET", "/graphql", nil)
	req.Header.Add("Authorization", "Bearer ey28718921th")

	if err != nil {
		t.Fatal(err)
	}

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwt, ok := r.Context().Value("jwt").(string)

		if !ok {
			t.Errorf("jwt not in request context: got %q", jwt)
		}

		if len(jwt) == 0 {
			t.Errorf("jwt should not be empty when it's passed")
		}
	})

	rr := httptest.NewRecorder()
	handler := server.PassJwtContext(testHandler)
	handler.ServeHTTP(rr, req)
}

// Don't provide a Authorization header, JWT should be a empty string since we
// pass the context to the GraphQL resolver func that handles authentication.
func TestJwtMiddleware2(t *testing.T) {
	req, err := http.NewRequest("GET", "/graphql", nil)

	if err != nil {
		t.Fatal(err)
	}

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwt, ok := r.Context().Value("jwt").(string)

		if !ok {
			t.Errorf("jwt not in request context: got %q", jwt)
		}

		if len(jwt) != 0 {
			t.Errorf("jwt should be empty when it's not passed")
		}
	})

	rr := httptest.NewRecorder()
	handler := server.PassJwtContext(testHandler)
	handler.ServeHTTP(rr, req)
}

func TestAuthMiddlewareValidToken(t *testing.T) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString(server.JwtSecret)
	callback := func() (interface{}, error) {
		return ResponseMock{"Doge"}, nil
	}

	if err != nil {
		t.Error(err)
	}

	res, error := server.AuthMiddleware(tokenString, callback)

	if error != nil {
		t.Error(error)
	}

	if diff := deep.Equal(res, ResponseMock{"Doge"}); diff != nil {
		t.Error(diff)
	}
}
