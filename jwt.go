package goutils

import (
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

// VerifyToken parses a jwt token
func VerifyToken(tokenString string, secretBytes []byte) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretBytes, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

// GenerateToken generates a jwt token string with provied claims
func GenerateToken(claims jwt.MapClaims, secretBytes []byte) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString(secretBytes)
	return signedToken
}
