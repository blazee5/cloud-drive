package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

const (
	signingKey = "AfdJ!@hj1#$#jhskFJFSkdfl"
)

type tokenClaims struct {
	jwt.RegisteredClaims
	UserID string `json:"user_id"`
}

func ParseToken(token string) (string, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := parsedToken.Claims.(*tokenClaims)
	if !ok {
		return "", err
	}

	return claims.UserID, nil
}
