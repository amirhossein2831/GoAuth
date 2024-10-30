package controller

import (
	"GoAuth/src/api/request/auth"
	"GoAuth/src/api/request/user"
	"GoAuth/src/pkg/response"
	"GoAuth/src/pkg/utils"
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

func (controller *AuthController) Refresh(c *gin.Context) response.IResponse {
	var req *auth.RefreshTokenRequest
	err := c.ShouldBind(&req)
	if err != nil {
		return response.NewResponse(c).SetError(err)
	}

	ctx := context.WithValue(context.Background(), "req", req)

	res, err := controller.Service.RefreshToken(ctx)
	if err != nil {
		return response.NewResponse(c).SetError(err)
	}

	return response.NewResponse(c).SetStatusCode(http.StatusOK).SetData(
		map[string]interface{}{
			"token": res,
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

func (controller *AuthController) Register(c *gin.Context) response.IResponse {
	var req *user.CreateUserRequest
	if err := c.ShouldBind(&req); err != nil {
		return response.NewResponse(c)
	}

	if err := val.Validate(req); err != nil {
		return response.NewResponse(c).SetData(err)
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

func (controller *AuthController) Update(c *gin.Context) response.IResponse {
	id, err := utils.GetID(c.Param("id"))
	if err != nil {
		return response.NewResponse(c).SetError(err)
	}

	var req *user.UpdateUserRequest
	err = c.ShouldBind(&req)
	if err != nil {
		return response.NewResponse(c)
	}

	if res := val.Validate(req); res != nil {
		return response.NewResponse(c).SetData(res)
	}

	ctx := context.WithValue(context.Background(), "req", req)
	ctx = context.WithValue(ctx, "columns", map[string]any{"id": id})
	res, err := controller.Service.Update(ctx)
	if err != nil {
		return response.NewResponse(c).SetError(err)
	}

	return response.NewResponse(c).SetStatusCode(http.StatusOK).
		SetData(map[string]any{
			"user": res,
		})
}

func (controller *AuthController) Verify(c *gin.Context) response.IResponse {
	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		return response.NewResponse(c).SetError(TokenIsMissed)
	}

	ctx := context.WithValue(context.Background(), "token", strings.TrimPrefix(accessToken, "Bearer "))

	res, err := controller.Service.Verify(ctx)
	if err != nil {
		return response.NewResponse(c).SetStatusCode(http.StatusUnprocessableEntity).SetError(err)
	}

	return response.NewResponse(c).SetStatusCode(http.StatusOK).SetData(map[string]interface{}{
		"is_valid": true,
		"claims":   res,
	})
}

func (controller *AuthController) Profile(c *gin.Context) response.IResponse {
	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		return response.NewResponse(c).SetError(TokenIsMissed)
	}

	ctx := context.WithValue(context.Background(), "token", strings.TrimPrefix(accessToken, "Bearer "))

	res, err := controller.Service.Profile(ctx)
	if err != nil {
		return response.NewResponse(c).SetStatusCode(http.StatusUnprocessableEntity).SetError(err)
	}

	return response.NewResponse(c).SetStatusCode(http.StatusOK).SetData(map[string]interface{}{
		"user": res,
	})
}

func (controller *AuthController) ChangePassword(c *gin.Context) response.IResponse {
	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		return response.NewResponse(c).SetError(TokenIsMissed)
	}

	var req *auth.ChangePasswordRequest
	if err := c.ShouldBind(&req); err != nil {
		return response.NewResponse(c)
	}

	if err := val.Validate(req); err != nil {
		return response.NewResponse(c).SetData(err)
	}

	ctx := context.WithValue(context.Background(), "req", req)
	ctx = context.WithValue(ctx, "token", strings.TrimPrefix(accessToken, "Bearer "))

	err := controller.Service.ChangePassword(ctx)
	if err != nil {
		return response.NewResponse(c).SetStatusCode(http.StatusUnprocessableEntity).SetError(err)
	}

	return response.NewResponse(c).SetStatusCode(http.StatusNoContent)
}

func (controller *AuthController) TokenList(c *gin.Context) response.IResponse {
	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		return response.NewResponse(c).SetError(TokenIsMissed)
	}

	ctx := context.WithValue(context.Background(), "token", strings.TrimPrefix(accessToken, "Bearer "))

	res, err := controller.Service.TokenList(ctx)
	if err != nil {
		return response.NewResponse(c).SetStatusCode(http.StatusUnprocessableEntity).SetError(err)
	}

	return response.NewResponse(c).SetStatusCode(http.StatusOK).SetData(map[string]any{
		"tokens": res,
	})
}
