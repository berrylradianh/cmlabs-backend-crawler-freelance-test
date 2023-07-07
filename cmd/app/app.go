package app

import (
	"github.com/berrylradianh/cmlabs-backend-crawler-freelance-test/internal/delivery"
	"github.com/berrylradianh/cmlabs-backend-crawler-freelance-test/internal/delivery/http/handler"
	authHandler "github.com/berrylradianh/cmlabs-backend-crawler-freelance-test/internal/delivery/http/handler/auth"
	userRepo "github.com/berrylradianh/cmlabs-backend-crawler-freelance-test/internal/domain/repository/user"
	authUsecase "github.com/berrylradianh/cmlabs-backend-crawler-freelance-test/internal/domain/usecase/auth"

	"github.com/berrylradianh/cmlabs-backend-crawler-freelance-test/pkg/database"
	"github.com/labstack/echo/v4"
)

func StartApp() *echo.Echo {
	database.Init()
	authRepo := userRepo.New(database.DB)
	authUsecase := authUsecase.New(authRepo)
	authHandler := authHandler.New(authUsecase)

	handler := handler.Handler{
		AuthHandler: authHandler,
	}

	router := delivery.StartRoute(handler)
	return router
}
