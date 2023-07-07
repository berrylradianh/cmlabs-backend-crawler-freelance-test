package crawling

import (
	"github.com/labstack/echo/v4"
)

func (ch *CrawlingHandler) RegisterRoutes(e *echo.Echo) {
	crawlingGroup := e.Group("/crawling")
	crawlingGroup.POST("/", ch.CrawlingWebsite())
}
