package crawling

import (
	cu "github.com/berrylradianh/cmlabs-backend-crawler-freelance-test/internal/domain/usecase/crawling"
)

type CrawlingHandler struct {
	crawlingUsecase cu.CrawlingUsecase
}

func NewCrawlingHandler(crawlingUsecase cu.CrawlingUsecase) *CrawlingHandler {
	return &CrawlingHandler{
		crawlingUsecase: crawlingUsecase,
	}
}
