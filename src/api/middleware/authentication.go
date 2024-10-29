package middleware

import (
	"GoAuth/src/models"
	"GoAuth/src/services"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AuthenticationMiddleware struct {
	AuthService services.IAuthService
}

func NewAuthenticationMiddleware() *AuthenticationMiddleware {
	return &AuthenticationMiddleware{
		AuthService: services.NewAuthService(),
	}
}

func (service *AuthenticationMiddleware) Middleware(userType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader("Authorization")
		if accessToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		ctx := context.WithValue(context.Background(), "token", strings.TrimPrefix(accessToken, "Bearer "))

		res, err := service.AuthService.Profile(ctx)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is invalid"})
			c.Abort()
			return

		}
		user := res.(*models.User)

		if string(user.Type) != userType {
			c.JSON(http.StatusForbidden, gin.H{"error": "You don't Have Permission"})
			c.Abort()
			return
		}

		c.Next()
	}

}
