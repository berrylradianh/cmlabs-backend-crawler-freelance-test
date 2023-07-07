package crawling

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"
)

func (cs *CrawlingService) SaveCSS(url string, data []byte) (string, error) {
	filename, err := GetFilename(url, ".css")
	if err != nil {
		return "", err
	}

	filePath := path.Join(".", filename)

	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		return "", fmt.Errorf("failed to save CSS file: %v", err)
	}

	return filePath, nil
}

func (cs *CrawlingService) SaveHTML(url string, data []byte) (string, error) {
	// log.Println(url)
	filename, err := GetFilename(url, ".html")
	if err != nil {
		return "", err
	}

	// log.Println(url)

	filePath := path.Join(".", filename)

	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		return "", fmt.Errorf("failed to save HTML file: %v", err)
	}

	return filePath, nil
}

func GetFilename(url string, extension string) (string, error) {
	filename := strings.ReplaceAll(url, "://", "-")
	filename = strings.ReplaceAll(filename, ".", "-")

	return fmt.Sprintf("%s%s", filename, extension), nil
}
