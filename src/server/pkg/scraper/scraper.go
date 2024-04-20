package scraper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var linkCache = make(map[string][]string)

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

func WriteCsv(filename string) error {
	file, err := os.Create("data.csv")
	if err != nil {
		return nil
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for key, links := range linkCache {
		row := append([]string{key}, links...)
		if err := writer.Write(row); err != nil {
			return err
		}

	}
	return nil

}

func ReadCsv(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Assuming there is no header row
	for {
		row, err := reader.Read()
		if err != nil {
			if err == csv.ErrFieldCount || strings.Contains(err.Error(), "EOF") {
				break // End of file or a line with wrong field count
			}
			return err
		}
		if len(row) > 1 {
			key := row[0]
			links := row[1:]
			linkCache[key] = links
		}
	}

	return nil
}

func LoadCache() {
	err := ReadCsv("data.csv")
	if err != nil {
		err2 := WriteCsv("data.csv")
		if err2 != nil {
			log.Fatal(err2)
		}

	}
}

func GetScrapeLinks(link string) []string {
	links, exist := linkCache[link]
	if !exist {
		err := WebScraping(link, &links)
		if err != nil {
			return nil
		}
		linkCache[link] = links
		return links
	}
	return links
}
