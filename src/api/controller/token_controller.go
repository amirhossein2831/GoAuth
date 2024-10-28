package controller

import (
	"GoAuth/src/pkg/response"
	"GoAuth/src/pkg/utils"
	"GoAuth/src/services"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TokenController struct {
	Service services.ITokenService
}

func NewTokenController() *TokenController {
	return &TokenController{
		Service: services.NewTokenService(),
	}
}

func (controller *TokenController) List(c *gin.Context) response.IResponse {
	ctx := context.Background()
	res, err := controller.Service.List(ctx)
	if err != nil {
		return response.NewResponse(c).SetError(err)
	}

	return response.NewResponse(c).SetStatusCode(http.StatusOK).
		SetData(map[string]any{
			"tokens": res,
		})
}

func (controller *TokenController) Get(c *gin.Context) response.IResponse {
	id, err := utils.GetID(c.Param("id"))
	if err != nil {
		return response.NewResponse(c).SetError(err)
	}

	ctx := context.WithValue(context.Background(), "tokenId", id)
	res, err := controller.Service.Get(ctx)
	if err != nil {
		return response.NewResponse(c).SetError(err)
	}

	return response.NewResponse(c).SetStatusCode(http.StatusOK).
		SetData(map[string]any{
			"token": res,
		})
}

func (controller *TokenController) Delete(c *gin.Context) response.IResponse {
	id, err := utils.GetID(c.Param("id"))
	if err != nil {
		return response.NewResponse(c).SetError(err)
	}

	ctx := context.WithValue(context.Background(), "tokenId", id)
	err = controller.Service.Delete(ctx)
	if err != nil {
		return response.NewResponse(c).SetError(err)
	}

	return response.NewResponse(c).SetStatusCode(http.StatusNoContent)
}
