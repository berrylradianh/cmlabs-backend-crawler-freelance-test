package handler

import (
	crawlingHandler "github.com/berrylradianh/cmlabs-backend-crawler-freelance-test/internal/delivery/http/handler/crawling"
)

type Handler struct {
	CrawlingHandler *crawlingHandler.CrawlingHandler
}
