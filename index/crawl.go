package index

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/gocolly/colly"
)

const PAGES_DIR = "pages"
const DEPTH = 1

const (
	HTTPS_SEP       = "//"
	BACKSLASH_SEP   = "/"
	UNDERSCORE      = "_"
	FILE_PERMISSION = 0644
)

const (
	HTML_TITLE = "title"
	HTML_URL   = "a[href]"
	HTML_TEXT  = "title, h1, h2, h3, h4, h5, h6, p, span, div"
)

func formatAllowedDomain(site string) string {
	allowedDomain := site
	parts := strings.Split(site, BACKSLASH_SEP)
	prefixExists := len(parts) > 2
	if prefixExists {
		allowedDomain = parts[2]
	}
	return allowedDomain
}

func urlToFileName(url string) string {
	parts := strings.Split(url, HTTPS_SEP)
	prefixExists := len(parts) > 1
	if prefixExists {
		url = parts[1]
	}
	return strings.ReplaceAll(url, BACKSLASH_SEP, UNDERSCORE)
}

func savePageText(url, content string) {
	if len(content) == 0 {
		return
	}

	filename := urlToFileName(url)
	filepath := filepath.Join(PAGES_DIR, filename)

	err := os.WriteFile(filepath, []byte(content), FILE_PERMISSION)
	if err != nil {
		fmt.Println("Error saving file:", err)
	} else {
		fmt.Println("Saved:", filepath)
	}
}

func crawlInner(currentUrl string, allowedDomain string, depth int) {
	if depth < 0 {
		return
	}

	collector := colly.NewCollector(colly.AllowedDomains(allowedDomain))

	var pageText string

	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Url", r.URL)
	})

	collector.OnError(func(e *colly.Response, err error) {
		fmt.Println("Request URL:", e.Request.URL, "failed with response:", e, "\nError:", err)
	})

	collector.OnHTML(HTML_TITLE, func(element *colly.HTMLElement) {
		fmt.Println("Title", element.Text)
	})

	collector.OnHTML(HTML_URL, func(element *colly.HTMLElement) {
		link := element.Request.AbsoluteURL(element.Attr("href"))
		crawlInner(link, allowedDomain, depth-1)
	})

	collector.OnHTML(HTML_TEXT, func(element *colly.HTMLElement) {
		pageText += element.Text + NEWLINE
	})

	collector.OnScraped(func(r *colly.Response) {
		savePageText(currentUrl, pageText)
	})

	err := collector.Visit(currentUrl)
	if err != nil {
		fmt.Println("Error visiting page:", err)
	}
}

func crawl(sites []string) {
	var wg sync.WaitGroup

	for _, site := range sites {
		wg.Add(1)
		go func(siteURL string) {
			defer wg.Done()
			allowedDomain := formatAllowedDomain(siteURL)

			crawlInner(siteURL, allowedDomain, DEPTH)
		}(site)
	}

	wg.Wait()
}
