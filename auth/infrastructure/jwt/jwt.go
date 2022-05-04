package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/middleware"
)

type jwtCustomClaims struct {
	UserID string `json:"userID"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

var signingKey = []byte("secret")

var JwtConfig = middleware.JWTConfig{
	Claims:     &jwtCustomClaims{},
	SigningKey: signingKey,
}

func NewToken(userID, email string) (string, error) {
	claims := &jwtCustomClaims{
		userID,
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signingKey)
}
