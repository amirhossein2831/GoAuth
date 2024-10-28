package routes

import (
	"GoAuth/src/api/controller"
	"GoAuth/src/pkg/net"
	"github.com/gin-gonic/gin"
)

func TokenRoutes(router *gin.RouterGroup) {
	tokenController := controller.NewTokenController()
	token := router.Group("tokens")

	token.GET("", net.HandlerFunc(tokenController.List))
	token.GET(":id", net.HandlerFunc(tokenController.Get))
	token.DELETE(":id", net.HandlerFunc(tokenController.Delete))
}
