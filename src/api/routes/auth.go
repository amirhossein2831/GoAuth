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
	auth.POST("register", net.HandlerFunc(authController.Register))
	auth.GET("logout", net.HandlerFunc(authController.Logout))
	auth.GET("verify", net.HandlerFunc(authController.Verify))
	auth.GET("profile", net.HandlerFunc(authController.Profile))
	auth.POST("change-password", net.HandlerFunc(authController.ChangePassword))
}
