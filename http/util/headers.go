package util

import (
	"errors"
	"net/http"
	"strings"
)

// Errors
var (
	ErrMissingAuthorization = errors.New("Missing Authorization")
	ErrInvalidAuthorization = errors.New("Invalid Authorization header")
)

// ParseBearerToken parses the JWT token from the request's
// Authorization: Bearer header
func ParseBearerToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", ErrMissingAuthorization
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		return "", ErrInvalidAuthorization
	}

	if headerParts[0] != "Bearer" {
		return "", ErrInvalidAuthorization
	}

	token := headerParts[1]

	return token, nil
}
