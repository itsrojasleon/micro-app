package internal

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	// "github.com/rojasleon/reserve-micro/auth/models"
)

type authClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

var jwtSecret = os.Getenv("JWT_SECRET")

// Generate JWT
// We only need the user's email to create a token.
func GenerateJWT(email string) (string, error) {
	claims := authClaims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    "auth-service",
			IssuedAt:  time.Now().Unix(),
		},
	}

	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := unsignedToken.SignedString([]byte(jwtSecret))

	if err != nil {
		return err.Error(), err
	}
	return signedToken, nil
}

// Verify JWT
func VerifyJWT(receivedToken string) jwt.MapClaims {
	token, err := jwt.Parse(receivedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", errors.New("unexpected signing method")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims
	}
	return nil
}
