package handler

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jmramos02/healthy-buddy/internal/domain/meal"
)

func CreateMealPlan(c *gin.Context) {
	ctx := c.MustGet("context").(context.Context)
	var request meal.CreateMealPlanRequest
	c.Bind(&request)

	response, err := meal.CreateMealPlan(ctx, request)

	render(response, err, c)
}
