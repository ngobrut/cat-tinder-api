package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	UserID string
	jwt.RegisteredClaims
}

const (
	JWT_TTL time.Duration = 8 * time.Hour
)

func GenerateAccessToken(claims *CustomClaims, secret string) (string, error) {
	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(JWT_TTL)),
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS512, claims).SignedString([]byte(secret))
}
