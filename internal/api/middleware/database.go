package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetDatabase(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.MustGet("context").(context.Context)
		ctx = context.WithValue(ctx, "db", db)

		c.Set("context", ctx)
		c.Next()
	}
}
