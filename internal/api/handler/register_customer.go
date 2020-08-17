package handler

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jmramos02/healthy-buddy/internal/features/authentication"
)

func RegisterCustomer(c *gin.Context) {
	var request authentication.RegisterRequest
	c.Bind(&request)

	ctx := c.MustGet("context").(context.Context)
	response, err := authentication.RegisterCustomer(ctx, request, "customer")

	render(response, err, c)
}
