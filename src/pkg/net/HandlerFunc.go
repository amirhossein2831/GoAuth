package net

import (
	"GoAuth/src/pkg/response"
	"github.com/gin-gonic/gin"
)

// HandlerFunc wraps a Gin handler to use the Response struct
func HandlerFunc(handler func(*gin.Context) response.IResponse) gin.HandlerFunc {
	return func(c *gin.Context) {
		res := handler(c)
		res.Send()
	}
}
