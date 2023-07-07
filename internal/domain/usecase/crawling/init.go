package crawling

import (
	"github.com/berrylradianh/cmlabs-backend-crawler-freelance-test/internal/domain/repository/crawling"
)

type CrawlingUsecase interface {
	SaveCSS(url string, data []byte) (string, error)
	SaveHTML(url string, data []byte) (string, error)
}

type crawlingUsecase struct {
	crawlingRepository crawling.CrawlingRepository
}

func NewCrawlingUsecase(crawlingRepository crawling.CrawlingRepository) CrawlingUsecase {
	return &crawlingUsecase{
		crawlingRepository: crawlingRepository,
	}
}
