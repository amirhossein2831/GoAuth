package routes

import (
	"GoAuth/src/api/controller"
	"GoAuth/src/api/middleware"
	"GoAuth/src/models"
	"GoAuth/src/pkg/net"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup) {
	userController := controller.NewUserController()
	users := router.Group("users").Use(middleware.NewAuthenticationMiddleware().Middleware(models.SuperAdmin))

	users.GET("", net.HandlerFunc(userController.List))
	users.GET(":id", net.HandlerFunc(userController.Get))
	users.POST("", net.HandlerFunc(userController.Create))
	users.PATCH(":id", net.HandlerFunc(userController.Update))
	users.DELETE(":id", net.HandlerFunc(userController.Delete))
	users.POST("change-password/:id", net.HandlerFunc(userController.ChangePassword))
}
