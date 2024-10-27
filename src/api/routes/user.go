package routes

import (
	"GoAuth/src/api/controller"
	"GoAuth/src/pkg/net"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup) {
	userController := controller.NewUserController()
	users := router.Group("users")

	users.GET("", net.HandlerFunc(userController.List))
	users.GET(":id", net.HandlerFunc(userController.Get))
	users.POST("", net.HandlerFunc(userController.Create))
	users.DELETE(":id", net.HandlerFunc(userController.Delete))
}
