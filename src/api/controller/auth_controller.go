package controller

import (
	"GoAuth/src/api/request/auth"
	"GoAuth/src/api/request/user"
	"GoAuth/src/pkg/response"
	val "GoAuth/src/pkg/validator"
	"GoAuth/src/services"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var TokenIsMissed = errors.New("token is not given")

type AuthController struct {
	Service services.IAuthService
}

func NewAuthController() *AuthController {
	return &AuthController{
		Service: services.NewAuthService(),
	}
}

func (controller *AuthController) Login(c *gin.Context) response.IResponse {
	var req *auth.LoginRequest
	err := c.ShouldBind(&req)
	if err != nil {
		return response.NewResponse(c).SetError(err)
	}

	ctx := context.WithValue(context.Background(), "req", req)
	res, err := controller.Service.Login(ctx)
	if err != nil {
		return response.NewResponse(c).SetError(err)
	}

	return response.NewResponse(c).SetStatusCode(http.StatusOK).SetData(
		map[string]interface{}{
			"token": res,
		})
}

func (controller *AuthController) Register(c *gin.Context) response.IResponse {
	var req *user.CreateUserRequest
	if err := c.ShouldBind(&req); err != nil {
		return response.NewResponse(c)
	}

	if res := val.Validate(req); res != nil {
		return response.NewResponse(c).SetData(res)
	}

	ctx := context.WithValue(context.Background(), "req", req)
	res, err := controller.Service.Register(ctx)
	if err != nil {
		return response.NewResponse(c).SetStatusCode(http.StatusUnprocessableEntity).SetError(err)
	}

	return response.NewResponse(c).SetStatusCode(http.StatusOK).
		SetData(map[string]any{
			"user": res,
		})
}

func (controller *AuthController) Logout(c *gin.Context) response.IResponse {
	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		return response.NewResponse(c).SetError(TokenIsMissed)
	}

	ctx := context.WithValue(context.Background(), "token", strings.TrimPrefix(accessToken, "Bearer "))
	err := controller.Service.Logout(ctx)
	if err != nil {
		return response.NewResponse(c).SetStatusCode(http.StatusUnprocessableEntity).SetError(err)
	}

	return response.NewResponse(c).SetStatusCode(http.StatusNoContent)
}
