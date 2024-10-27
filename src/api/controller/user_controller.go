package controller

import (
	"GoAuth/src/api/request/user"
	"GoAuth/src/pkg/response"
	val "GoAuth/src/pkg/validator"
	"GoAuth/src/services"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	Service services.IUserService
}

func NewUserController() *UserController {
	return &UserController{
		Service: services.NewUserService(),
	}
}

func (controller *UserController) List(c *gin.Context) response.IResponse {
	ctx := context.Background()
	res, err := controller.Service.List(ctx)
	if err != nil {
		return response.NewResponse(c).SetMessage("Cannot get List of users")
	}

	return response.NewResponse(c).SetStatusCode(http.StatusOK).
		SetData(map[string]any{
			"users": res,
		})
}

func (controller *UserController) Get(c *gin.Context) response.IResponse {
	ctx := context.WithValue(context.Background(), "userId", c.Param("id"))
	res, err := controller.Service.Get(ctx)
	if err != nil {
		return response.NewResponse(c).SetMessage("Cannot get user")
	}

	return response.NewResponse(c).SetStatusCode(http.StatusOK).
		SetData(map[string]any{
			"user": res,
		})
}

func (controller *UserController) Create(c *gin.Context) response.IResponse {
	var req *user.CreateUserRequest
	if err := c.ShouldBind(&req); err != nil {
		return response.NewResponse(c)
	}

	if res := val.Validate(req); res != nil {
		return response.NewResponse(c).SetData(res)
	}

	ctx := context.WithValue(context.Background(), "req", req)
	res, err := controller.Service.Create(ctx)
	if err != nil {
		return response.NewResponse(c).SetStatusCode(http.StatusUnprocessableEntity).SetMessage("Cannot create user")
	}

	return response.NewResponse(c).SetStatusCode(http.StatusOK).
		SetData(map[string]any{
			"user": res,
		})
}

func (controller *UserController) Delete(c *gin.Context) response.IResponse {
	ctx := context.WithValue(context.Background(), "userId", c.Param("id"))
	err := controller.Service.Delete(ctx)
	if err != nil && errors.Is(err, services.UserNotFound) {
		return response.NewResponse(c).SetMessage(err.Error())
	}

	if err != nil {
		return response.NewResponse(c).SetStatusCode(http.StatusUnprocessableEntity).SetMessage("Cannot delete user")
	}

	return response.NewResponse(c).SetStatusCode(http.StatusNoContent)
}
