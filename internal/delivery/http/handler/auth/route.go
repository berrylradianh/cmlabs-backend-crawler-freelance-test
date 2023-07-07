package auth

import (
	"github.com/labstack/echo/v4"
)

func (authHandler *AuthHandler) RegisterRoutes(e *echo.Echo) {
	authGroup := e.Group("/account")
	authGroup.POST("/register", authHandler.Register())
	authGroup.POST("/login", authHandler.Login())
}
