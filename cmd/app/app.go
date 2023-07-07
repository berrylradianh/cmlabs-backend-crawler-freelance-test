package app

import (
	"github.com/berrylradianh/cmlabs-backend-crawler-freelance-test/internal/delivery"
	"github.com/berrylradianh/cmlabs-backend-crawler-freelance-test/internal/delivery/http/handler"

	crawlingHandler "github.com/berrylradianh/cmlabs-backend-crawler-freelance-test/internal/delivery/http/handler/crawling"
	crawlingRepo "github.com/berrylradianh/cmlabs-backend-crawler-freelance-test/internal/domain/repository/crawling"
	crawlingUsecase "github.com/berrylradianh/cmlabs-backend-crawler-freelance-test/internal/domain/usecase/crawling"

	"github.com/labstack/echo/v4"
)

func StartApp() *echo.Echo {
	crawlingRepo := crawlingRepo.New()
	crawlingUsecase := crawlingUsecase.NewCrawlingUsecase(crawlingRepo)
	crawlingHandler := crawlingHandler.NewCrawlingHandler(crawlingUsecase)

	handler := handler.Handler{
		CrawlingHandler: crawlingHandler,
	}

	router := delivery.StartRoute(handler)
	return router
}
