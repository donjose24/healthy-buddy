package middleware

import (
	"context"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jmramos02/healthy-buddy/internal/model"
)

type Claim struct {
	User    model.User
	Details interface{}
	jwt.StandardClaims
}

func ValidateUserSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqToken := c.Request.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer")
		if len(splitToken) != 2 {
			c.JSON(401, gin.H{
				"code": "UNAUTHORIZED",
			})
			c.Abort()
			return
		}

		reqToken = strings.TrimSpace(splitToken[1])
		appKey := os.Getenv("APPLICATION_KEY")
		claims := &Claim{}

		_, err := jwt.ParseWithClaims(reqToken, claims, func(token *jwt.Token) (interface{}, error) {
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

		c.Set("context", ctx)
		c.Next()
	}
}
