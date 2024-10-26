package api

import (
	"GoAuth/src/api/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func Init() error {
	return initServer()
}

func getNewRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	// Initialize new app.
	router := gin.New()

	// Attach Global Middleware here
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	return router
}

func initServer() error {
	// Initial v1 routes
	router := getNewRouter()
	v1 := router.Group("api/v1")
	{
		routes.UserRoutes(v1)
	}

	// Run App.
	log.Println("API Service: Initialized Successfully.")
	if err := router.Run(
		fmt.Sprintf("%s:%s",
			os.Getenv("APP_HOST"),
			os.Getenv("APP_PORT"),
		),
	); err != nil {
		return err
	}

	return nil
}
