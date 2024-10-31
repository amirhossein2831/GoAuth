package controller

import (
	"GoAuth/src/api/request/user"
	ctx2 "GoAuth/src/pkg/ctx"
	"GoAuth/src/pkg/response"
	"GoAuth/src/pkg/utils"
	val "GoAuth/src/pkg/validator"
	"GoAuth/src/services"
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
	ctx := ctx2.New().SetMap("columns", "", nil)
	res, err := controller.Service.List(ctx)
	if err != nil {
		return response.NewResponse(c).SetError(err)
	}

	return response.NewResponse(c).SetStatusCode(http.StatusOK).
		SetData(map[string]any{
			"users": res,
		})
}

func (controller *UserController) Get(c *gin.Context) response.IResponse {
	id, err := utils.GetID(c.Param("id"))
	if err != nil {
		return response.NewResponse(c).SetError(err)
	}

	ctx := ctx2.New().SetMap("columns", "id", id)
	res, err := controller.Service.Get(ctx)
	if err != nil {
		return response.NewResponse(c).SetError(err)
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

	ctx := ctx2.New().Set("req", req)
	res, err := controller.Service.Create(ctx)
	if err != nil {
		return response.NewResponse(c).SetStatusCode(http.StatusUnprocessableEntity).SetError(err)
	}

	return response.NewResponse(c).SetStatusCode(http.StatusOK).
		SetData(map[string]any{
			"user": res,
		})
}

func (controller *UserController) Update(c *gin.Context) response.IResponse {
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

	ctx := ctx2.New().SetMap("columns", "id", id).Set("req", req)
	res, err := controller.Service.Update(ctx)
	if err != nil {
		return response.NewResponse(c).SetError(err)
	}

	return response.NewResponse(c).SetStatusCode(http.StatusOK).
		SetData(map[string]any{
			"user": res,
		})
}

func (controller *UserController) Delete(c *gin.Context) response.IResponse {
	id, err := utils.GetID(c.Param("id"))
	if err != nil {
		return response.NewResponse(c).SetError(err)
	}

	ctx := ctx2.New().SetMap("columns", "id", id)
	err = controller.Service.Delete(ctx)
	if err != nil {
		return response.NewResponse(c).SetError(err)
	}

	return response.NewResponse(c).SetStatusCode(http.StatusNoContent)
}

func (controller *UserController) ChangePassword(c *gin.Context) response.IResponse {
	var req *user.ChangePasswordRequest
	if err := c.ShouldBind(&req); err != nil {
		return response.NewResponse(c)
	}

	if err := val.Validate(req); err != nil {
		return response.NewResponse(c).SetData(err)
	}

	id, err := utils.GetID(c.Param("id"))
	if err != nil {
		return response.NewResponse(c).SetError(err)
	}

	ctx := ctx2.New().SetMap("columns", "id", id).Set("req", req)
	_, err = controller.Service.ChangePassword(ctx)
	if err != nil {
		return response.NewResponse(c).SetError(err)
	}

	return response.NewResponse(c).SetStatusCode(http.StatusNoContent)
}
