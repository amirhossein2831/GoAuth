package routes

import (
	"GoAuth/src/api/controller"
	"GoAuth/src/pkg/net"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup) {
	userController := &controller.UserController{}
	router.GET("users", net.HandlerFunc(userController.List))
}
