package crawling

import (
	rc "github.com/berrylradianh/cmlabs-backend-crawler-freelance-test/internal/domain/repository/crawling"
)

type CrawlingUsecase interface {
	SaveCSS(url string, data []byte) (string, error)
	SaveHTML(url string, data []byte) (string, error)
	GetFilename(url string, extension string) (string, error)
}

type crawlingUsecase struct {
	crawlingRepository rc.CrawlingRepository
}

func New(crawlingRepository rc.CrawlingRepository) *crawlingUsecase {
	return &crawlingUsecase{
		crawlingRepository,
	}
}
