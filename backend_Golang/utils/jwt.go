package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nineard99/restaurant-Golang/types"
)

func GenerateJWT(userID string, role string) (string, error) {
	// TimeOut: 7 Day
	expiration := time.Now().Add(7 * 24 * time.Hour)

	claims := types.JWTClaims{
		ID:   userID,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// Read token และแปลงเป็น claims
func ParseJWT(tokenStr string) (*types.JWTClaims, error) {
	secret := os.Getenv("JWT_SECRET")
	
	token, err := jwt.ParseWithClaims(tokenStr, &types.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(*types.JWTClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
