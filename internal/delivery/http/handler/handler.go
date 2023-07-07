package handler

import (
	authHandler "github.com/berrylradianh/cmlabs-backend-crawler-freelance-test/internal/delivery/http/handler/auth"
	crawlingHandler "github.com/berrylradianh/cmlabs-backend-crawler-freelance-test/internal/delivery/http/handler/crawling"
)

type Handler struct {
	AuthHandler     *authHandler.AuthHandler
	CrawlingHandler *crawlingHandler.CrawlingHandler
}
