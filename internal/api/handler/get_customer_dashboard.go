package handler

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jmramos02/healthy-buddy/internal/domain/customer"
)

func GetCustomerDashboard(c *gin.Context) {
	ctx := c.MustGet("context").(context.Context)
	response, err := customer.Fetch(ctx)

	render(response, err, c)
}
