package controller

import (
	"GoAuth/src/api/request/auth"
	"GoAuth/src/pkg/response"
	"GoAuth/src/services"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	Service services.IAuthService
}

func NewAuthController() *AuthController {
	return &AuthController{
		Service: services.NewAuthService(),
	}
}

func (authController *AuthController) Login(c *gin.Context) response.IResponse {
	var req *auth.LoginRequest
	err := c.ShouldBind(&req)
	if err != nil {
		return response.NewResponse(c).SetMessage(err.Error())
	}

	ctx := context.WithValue(context.Background(), "req", req)
	res, err := authController.Service.Login(ctx)
	if err != nil {
		return response.NewResponse(c).SetMessage(err.Error())
	}

	return response.NewResponse(c).SetStatusCode(http.StatusOK).SetData(
		map[string]interface{}{
			"token": res,
		})
}

//func (authController *AuthController) Register(c *gin.Context) response.Response {}
