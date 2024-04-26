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
	"sync"
	"time"
	"unicode"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

var LinkCache = make(map[string][]string)
var CounterArticle = 0
var mutexCache = sync.Mutex{}

func WebScraping(url string, resultData *[]string) error {

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("error making GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
	}

	hasSeen := make(map[string]bool)
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err
	}

	doc.Find("#bodyContent a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists && strings.HasPrefix(href, "/wiki/") && !strings.HasPrefix(href, "/wiki/File:") && !hasSeen[href] {
			*resultData = append(*resultData, "https://en.wikipedia.org"+href)
			hasSeen[href] = true
		}
	})

	return nil
}

func GetScrapeLinksConcurrent(link string) []string {
	mutexCache.Lock()
	links, exist := LinkCache[link]
	mutexCache.Unlock()

	if !exist {
		links = []string{}
		if err := WebScraping(link, &links); err != nil {
			return nil
		}

		mutexCache.Lock()
		LinkCache[link] = links
		CounterArticle += 1
		mutexCache.Unlock()
	}
	return links
}

func WriteCsv(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return nil
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for key, links := range LinkCache {
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
	err = encoder.Encode(LinkCache)
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
			LinkCache[key] = links
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
	err = decoder.Decode(&LinkCache)
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

// IsWikiPageUrlExists checks if a Wikipedia page URL exists and updates it to the redirected URL.
func IsWikiPageUrlExists(url *string) bool {
	// Custom HTTP client with redirect policy
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// Update the original URL to the last requested URL
			*url = req.URL.String()
			return nil // Continue following redirects
		},
		Timeout: 10 * time.Second, // Set a timeout to avoid hanging on slow networks
	}

	// Perform the HTTP GET request
	response, err := client.Get(*url)
	if err != nil {
		fmt.Printf("HTTP request failed: %v\n", err)
		return false
	}

	defer response.Body.Close() // Ensure the response body is closed

	// Update the URL to the final redirected URL if not already done
	*url = response.Request.URL.String()

	// Check if the status code is in the 2xx range, indicating success
	return response.StatusCode >= 200 && response.StatusCode < 300
}

func LoadCache() {
	CounterArticle = 0
	err := ReadJSON("links.json")
	if err != nil {
		err2 := WriteJSON("links.json")
		if err2 != nil {
			log.Fatal(err2)
		}

	}
}

func WebScrapingColly(url string, resultData *[]string) error {
	if !IsWikiPageUrlExists(&url) {
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
	links, exist := LinkCache[link]
	if !exist {
		err := WebScraping(link, &links)
		if err != nil {
			return nil
		}
		if links != nil {

			LinkCache[link] = links
		}

		return links
	}
	return links
}

func GetScrapeLinksColly(link string) []string {
	links, exist := LinkCache[link]
	if !exist {
		err := WebScrapingColly(link, &links)
		if err != nil {
			return nil
		}
		LinkCache[link] = links
		return links
	}
	return links
}
