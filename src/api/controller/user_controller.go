package controller

import (
	"GoAuth/src/api/request/user"
	"GoAuth/src/pkg/response"
	"GoAuth/src/pkg/utils"
	val "GoAuth/src/pkg/validator"
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
		return response.NewResponse(c).SetError(err.Error())
	}

	return response.NewResponse(c).SetStatusCode(http.StatusOK).
		SetData(map[string]any{
			"users": res,
		})
}

func (controller *UserController) Get(c *gin.Context) response.IResponse {
	id, err := utils.GetID(c.Param("id"))
	if err != nil {
		return response.NewResponse(c).SetError(err.Error())
	}

	ctx := context.WithValue(context.Background(), "userId", id)
	res, err := controller.Service.Get(ctx)
	if err != nil {
		return response.NewResponse(c).SetError(err.Error())
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
		return response.NewResponse(c).SetStatusCode(http.StatusUnprocessableEntity).SetError(err.Error())
	}

	return response.NewResponse(c).SetStatusCode(http.StatusOK).
		SetData(map[string]any{
			"user": res,
		})
}

func (controller *UserController) Update(c *gin.Context) response.IResponse {
	id, err := utils.GetID(c.Param("id"))
	if err != nil {
		return response.NewResponse(c).SetError(err.Error())
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
	ctx = context.WithValue(ctx, "userId", id)
	res, err := controller.Service.Update(ctx)
	if err != nil {
		return response.NewResponse(c).SetError(err.Error())
	}

	return response.NewResponse(c).SetStatusCode(http.StatusOK).
		SetData(map[string]any{
			"user": res,
		})
}

func (controller *UserController) Delete(c *gin.Context) response.IResponse {
	id, err := utils.GetID(c.Param("id"))
	if err != nil {
		return response.NewResponse(c).SetError(err.Error())
	}

	ctx := context.WithValue(context.Background(), "userId", id)
	err = controller.Service.Delete(ctx)
	if err != nil {
		return response.NewResponse(c).SetError(err.Error())
	}

	return response.NewResponse(c).SetStatusCode(http.StatusNoContent)
}
