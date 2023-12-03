package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"math/big"
	"time"
)

const (
	signingKey = "AfdJ!@hj1#$#jhskFJFSkdfl"
	tokenTTL   = 12 * time.Hour
	salt       = "kaSDFklj$fds@#"
	codeLength = 8
)

type tokenClaims struct {
	jwt.RegisteredClaims
	UserID string `json:"user_id"`
}

func GenerateToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		userID,
	})

	return token.SignedString([]byte(signingKey))
}

func GenerateHashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func GenerateRandomCode() (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result string

	for i := 0; i < codeLength; i++ {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result += string(charset[randomIndex.Int64()])
	}

	return result, nil
}
