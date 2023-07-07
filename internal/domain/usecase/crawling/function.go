package crawling

func (uc *crawlingUsecase) SaveCSS(url string, data []byte) (string, error) {
	return uc.crawlingRepository.SaveCSS(url, data)
}

func (uc *crawlingUsecase) SaveHTML(url string, data []byte) (string, error) {
	return uc.crawlingRepository.SaveHTML(url, data)
}
