package models

import "github.com/golang-jwt/jwt"

type Claims struct {
	Uuid  string
	Email string
	Role  string
	jwt.StandardClaims
}
