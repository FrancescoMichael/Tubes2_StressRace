package scraper

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func WebScraping(url string, resultData *[]string) error {
	// Check if url is wikipedia
	if !strings.Contains(url, "wikipedia.org") {
		return fmt.Errorf("invalid URL: only Wikipedia articles are allowed")
	}

	// use net/http to get and validate
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error making GET request: %v", err)
	}
	defer resp.Body.Close()

	// Check the HTTP status code
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
	}

	// Create a goquery document from the HTTP response body
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err
	}

	// Find and process all links in the #bodyContent element
	doc.Find("#bodyContent a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if !exists {
			return // Skip this link if the href attribute does not exist
		}
		// Check if the link is internal to Wikipedia
		if strings.HasPrefix(href, "/wiki/") {
			*resultData = append(*resultData, "https://en.wikipedia.org"+href)
		}
	})

	return nil
}
