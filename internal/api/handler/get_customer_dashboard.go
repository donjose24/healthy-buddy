package handler

import (
	"context"

	"github.com/gin-gonic/gin"
	customer "github.com/jmramos02/healthy-buddy/internal/features/customer_dashboard"
)

func GetCustomerDashboard(c *gin.Context) {
	ctx := c.MustGet("context").(context.Context)
	response, err := customer.Fetch(ctx)

	render(response, err, c)
}
