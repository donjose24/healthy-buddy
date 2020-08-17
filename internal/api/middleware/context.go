package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
)

func SetContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		appContext := context.TODO()
		c.Set("context", appContext)
		c.Next()
	}
}
