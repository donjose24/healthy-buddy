package middleware

import (
	"context"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jmramos02/healthy-buddy/internal/model"
)

type Claim struct {
	User    model.User
	Details interface{}
	jwt.StandardClaims
}

func ValidateUserSession(token string) gin.HandlerFunc {
	return func(c *gin.Context) {
		appKey := os.Getenv("APPLICATION_KEY")
		claims := &Claim{}

		_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(appKey), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"code": "UNAUTHORIZED",
			})
		}

		ctx := c.MustGet("context").(context.Context)
		ctx = context.WithValue(ctx, "user", claims.User)
		ctx = context.WithValue(ctx, "details", claims.Details)
	}
}
