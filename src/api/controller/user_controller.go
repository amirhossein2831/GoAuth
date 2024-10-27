package controller

import (
	"GoAuth/src/pkg/response"
	"GoAuth/src/services"
	"context"
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
		return response.NewResponse(c).SetStatusCode(http.StatusBadRequest).SetMessage("Cannot get List of users")
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
		return response.NewResponse(c).SetStatusCode(http.StatusBadRequest).SetMessage("Cannot get user")
	}

	return response.NewResponse(c).SetStatusCode(http.StatusOK).
		SetData(map[string]any{
			"user": res,
		})
}
