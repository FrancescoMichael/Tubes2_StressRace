package scraper

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var LinkCache = make(map[string][]string)
var mutexCache = sync.Mutex{}
var Unique map[string]bool

func WebScraping(url string, resultData *[]string) error {

	client := &http.Client{Timeout: 10 * time.Second} // timeout
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
		if exists && strings.HasPrefix(href, "/wiki/") && !strings.HasPrefix(href, "/wiki/File:") && !hasSeen[href] && !strings.HasPrefix(href, "/wiki/Category:") && !strings.HasPrefix(href, "/wiki/Template:") && !strings.HasPrefix(href, "/wiki/Special:") && !strings.HasPrefix(href, "/wiki/Wikipedia:") && !strings.HasPrefix(href, "/wiki/Help:") && !strings.HasPrefix(href, "/wiki/Portal:") && !strings.HasPrefix(href, "/wiki/Template_talk:"){
			*resultData = append(*resultData, "https://en.wikipedia.org"+href)
			hasSeen[href] = true
		}
	})

	return nil
}

func GetScrapeLinksConcurrent(link string) []string {
	mutexCache.Lock()
	Unique[link] = true
	links, exist := LinkCache[link]
	mutexCache.Unlock()

	if !exist {
		links = []string{}
		if err := WebScraping(link, &links); err != nil {
			return nil
		}

		mutexCache.Lock()
		LinkCache[link] = links
		mutexCache.Unlock()
	}
	return links
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
	Unique = make(map[string]bool)
	err := ReadJSON("links.json")
	if err != nil {
		err2 := WriteJSON("links.json")
		if err2 != nil {
			log.Fatal(err2)
		}

	}
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
