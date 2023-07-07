package crawling

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gocolly/colly/v2"

	ue "github.com/berrylradianh/cmlabs-backend-crawler-freelance-test/internal/domain/entity/url"

	"github.com/labstack/echo/v4"
)

func (ch *CrawlingHandler) CrawlingWebsite() echo.HandlerFunc {
	return func(e echo.Context) error {
		var request ue.Request
		var links []string

		if err := e.Bind(&request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"Message": err.Error(),
			})
		}

		for _, url := range request.Url {
			c := colly.NewCollector()

			c.OnHTML("a[href]", func(e *colly.HTMLElement) {
				link := e.Attr("href")
				if strings.HasPrefix(link, "/") {
					link = fmt.Sprintf("%s%s", url, link)
				}
				links = append(links, link)
			})

			c.OnResponse(func(r *colly.Response) {
				contentType := r.Headers.Get("Content-Type")

				if strings.Contains(contentType, "text/html") {
					ch.crawlingUsecase.SaveHTML(url, r.Body)
				}

				if strings.Contains(contentType, "text/css") {
					ch.crawlingUsecase.SaveCSS(url, r.Body)
				}
			})

			err := c.Visit(url)
			if err != nil {
				return e.JSON(http.StatusExpectationFailed, map[string]interface{}{
					"Message": "Failed to crawl " + url,
				})
			}

			links = append(links, "-----------------------------------------------------------------")
			links = append(links, "---------------------Website Crawl Finished----------------------")
			links = append(links, "-----------------------------------------------------------------")
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"Links":   links,
			"Message": "Website crawled successfully",
		})
	}
}
