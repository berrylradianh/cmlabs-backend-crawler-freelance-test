package crawling

type CrawlingRepository interface {
	SaveCSS(url string, data []byte) (string, error)
	SaveHTML(url string, data []byte) (string, error)
}

type CrawlingService struct {
	repo CrawlingRepository
}

func New() CrawlingRepository {
	return &CrawlingService{}
}
