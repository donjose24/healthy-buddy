package middleware

import (
	"context"
	"fmt"
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

func ValidateUserSession(userType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqToken := c.Request.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer")
		if len(splitToken) != 2 {
			fmt.Println(splitToken)
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
			fmt.Println(err.Error())
			c.AbortWithStatusJSON(401, gin.H{
				"code": "UNAUTHORIZED",
			})
			return
		}

		if claims.User.Type != userType {
			c.AbortWithStatusJSON(403, gin.H{
				"code": "FORBIDDEN",
			})
			return
		}

		ctx := c.MustGet("context").(context.Context)
		ctx = context.WithValue(ctx, "user", claims.User)
		ctx = context.WithValue(ctx, "details", claims.Details)

		c.Set("context", ctx)
		c.Next()
	}
}
