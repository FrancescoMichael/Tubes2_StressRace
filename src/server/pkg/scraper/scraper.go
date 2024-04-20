package scraper

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
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
func WriteJSON(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(linkCache)
	if err != nil {
		return err
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

func ReadJSON(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&linkCache)
	if err != nil {
		return err
	}

	return nil
}

func LoadCache() {
	err := ReadJSON("data.json")
	if err != nil {
		err2 := WriteJSON("data.json")
		if err2 != nil {
			log.Fatal(err2)
		}

	}
}

func WebScrapingColly(url string, resultData *[]string) error {
	if !strings.Contains(url, "wikipedia.org") {
		return fmt.Errorf("invalid URL: only Wikipedia articles are allowed")
	}

	c := colly.NewCollector(
		colly.AllowedDomains("wikipedia.org", "en.wikipedia.org"),
	)

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnHTML("#bodyContent a[href]", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		// Check if the link is internal to Wikipedia
		if strings.HasPrefix(href, "/wiki/") {
			completeLink := "https://en.wikipedia.org" + href
			*resultData = append(*resultData, completeLink)
		}
	})
	// Start scraping
	err := c.Visit(url)
	if err != nil {
		return err
	}
	c.Wait()
	return nil
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

func GetScrapeLinksColly(link string) []string {
	links, exist := linkCache[link]
	if !exist {
		err := WebScrapingColly(link, &links)
		if err != nil {
			return nil
		}
		linkCache[link] = links
		return links
	}
	return links
}
