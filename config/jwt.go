package config

import "github.com/golang-jwt/jwt/v5"

var JWT_KEY = []byte("secret")

type JWTUserData struct {
	Username string
	jwt.RegisteredClaims
}
