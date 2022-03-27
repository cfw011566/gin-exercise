package middlewares

import (
	"example/golang-gin-poc/service"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" && strings.Contains(authHeader, BEARER_SCHEMA) {
			tokenString := authHeader[len(BEARER_SCHEMA):]

			token, err := service.NewJWTService().ValidateToken(tokenString)

			if token != nil && token.Valid {
				claims := token.Claims.(jwt.MapClaims)
				log.Println("Claims[Name]: ", claims["name"])
				log.Println("Claims[Admin]: ", claims["admin"])
				log.Println("Claims[Issuer]: ", claims["iss"])
				log.Println("Claims[IssuedAt]: ", claims["iat"])
				log.Println("Claims[ExpiredAt]: ", claims["exp"])
			} else {
				log.Println(err)
				c.AbortWithStatus(http.StatusUnauthorized)
			}
		} else {
			log.Println("No Authorization Header")
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
