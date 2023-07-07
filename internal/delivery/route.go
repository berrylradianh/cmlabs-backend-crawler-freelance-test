package delivery

import (
	"github.com/berrylradianh/cmlabs-backend-crawler-freelance-test/internal/delivery/http/handler"
	hlp "github.com/berrylradianh/cmlabs-backend-crawler-freelance-test/internal/domain/helper"
	"github.com/berrylradianh/cmlabs-backend-crawler-freelance-test/pkg/logger"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func StartRoute(handler handler.Handler) *echo.Echo {
	e := echo.New()
	e.Validator = &hlp.CustomValidator{Validator: validator.New()}
	logger.LogMiddleware(e)

	handler.AuthHandler.RegisterRoutes(e)

	return e
}
