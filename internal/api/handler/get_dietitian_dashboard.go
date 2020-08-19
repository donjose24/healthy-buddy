package handler

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jmramos02/healthy-buddy/internal/domain/dietitian"
)

func GetDietitianDashboard(c *gin.Context) {
	ctx := c.MustGet("context").(context.Context)
	response, err := dietitian.Fetch(ctx)

	render(response, err, c)
}
