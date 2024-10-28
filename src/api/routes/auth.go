package routes

import (
	"GoAuth/src/api/controller"
	"GoAuth/src/pkg/net"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup) {
	authController := controller.NewAuthController()
	auth := router.Group("auth")

	auth.POST("login", net.HandlerFunc(authController.Login))
	//auth.POST(":id", net.HandlerFunc(authController.Get))
	auth.GET("logout", net.HandlerFunc(authController.Logout))
}
