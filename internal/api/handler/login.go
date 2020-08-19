package handler

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jmramos02/healthy-buddy/internal/domain/authentication"
)

func Login(c *gin.Context) {
	var request authentication.LoginRequest
	c.Bind(&request)

	ctx := c.MustGet("context").(context.Context)
	response, err := authentication.Login(ctx, request)

	render(response, err, c)
}
