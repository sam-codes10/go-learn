package middleware

import (
	"auth-service/models"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func GenerateToken(claims models.Claims) (string, error) {
	secretKey := os.Getenv("GO_JWT_SECRETKEY")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}
