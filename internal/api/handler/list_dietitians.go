package handler

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jmramos02/healthy-buddy/internal/domain/dietitian"
)

func ListDietitians(c *gin.Context) {
	ctx := c.MustGet("context").(context.Context)
	response, err := dietitian.List(ctx)

	render(response, err, c)
}
