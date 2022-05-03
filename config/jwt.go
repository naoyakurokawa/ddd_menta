package config

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/middleware"
)

type jwtCustomClaims struct {
	UserID string `json:"userID"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

var SigningKey = []byte("secret")

var JwtConfig = middleware.JWTConfig{
	Claims:     &jwtCustomClaims{},
	SigningKey: SigningKey,
}
