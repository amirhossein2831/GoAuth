package controller

import (
	"GoAuth/src/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
}

func (u UserController) List(c *gin.Context) response.IResponse {
	return response.NewResponse(c).SetStatusCode(http.StatusOK).
		SetData(map[string]any{
			"users": "think this is all users",
		})
}
