package scraper

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"unicode"

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
	file, err := os.Create(filename)
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

func TitleToWikiUrl(title string) string {
	title = strings.TrimSpace(title)
	title = toCamelCase(title)
	title = strings.ReplaceAll(title, " ", "_")
	title = url.QueryEscape(title)
	return fmt.Sprintf("https://en.wikipedia.org/wiki/%s", title)
}

func wikiUrlToTitle(wikiUrl string) string {
	decoded, err := url.QueryUnescape(strings.TrimPrefix(wikiUrl, "https://en.wikipedia.org/wiki/"))
	if err != nil {
		return ""
	}
	return strings.ReplaceAll(decoded, "_", " ")
}

func PathToTitle(path []string) []string {
	var hasil []string
	for _, link := range path {
		hasil = append(hasil, wikiUrlToTitle(link))
	}
	return hasil
}

// toCamelCase converts the title into title case. This is only experimental, user best write the title correctly !
func toCamelCase(input string) string {
	words := strings.Fields(input) // split string with white space delimiter
	for i, word := range words {
		runes := []rune(word)            // convert string to slice of runes, deal with string beyonc ascii
		if i == 0 || i == len(words)-1 { // Always capitalize the first and last word
			runes[0] = unicode.ToUpper(runes[0])
		} else {
			// Lower case for specific small words in the middle of a title
			lower := strings.ToLower(word)
			if lower == "the" || lower == "and" || lower == "in" || lower == "of" || lower == "a" {
				runes = []rune(lower)
			} else {
				runes[0] = unicode.ToUpper(runes[0])
			}
		}
		for j := 1; j < len(runes); j++ {
			runes[j] = unicode.ToLower(runes[j])
		}
		words[i] = string(runes)
	}
	return strings.Join(words, " ")
}

func IsWikiPageUrlExists(url string) bool {

	response, err := http.Get(url)
	if err != nil {
		return false
	}
	defer response.Body.Close()
	return response.StatusCode == 200
}

func LoadCache() {
	err := ReadJSON("links.json")
	if err != nil {
		err2 := WriteJSON("links.json")
		if err2 != nil {
			log.Fatal(err2)
		}

	}
}

func WebScrapingColly(url string, resultData *[]string) error {
	// if !strings.Contains(url, "wikipedia.org") {
	// 	return fmt.Errorf("invalid URL: only Wikipedia articles are allowed")
	// }
	if !IsWikiPageUrlExists(url) {
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
		if links != nil {
			linkCache[link] = links
		}

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
