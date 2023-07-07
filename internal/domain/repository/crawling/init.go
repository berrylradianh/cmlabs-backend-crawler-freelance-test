package crawling

type CrawlingRepository interface {
	SaveCSS(url string, data []byte) (string, error)
	SaveHTML(url string, data []byte) (string, error)
	GetFilename(url string, extension string) (string, error)
}

type crawlingRepo struct {
	// No Gorm  || No Database
}

func New() CrawlingRepository {
	return &crawlingRepo{}
}
