package scraper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var LinkCache = make(map[string][]string) // store cache links
var mutexCache = sync.Mutex{}             // mutex for accessing linkCache
var Unique map[string]bool                // store unique urls

// func WebScraping will scrape all unique urls
func WebScraping(url string, resultData *[]string) error {

	client := &http.Client{Timeout: 10 * time.Second} // timeout, 10 second
	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("error making GET request: %v", err) // getting request error
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-200 status code: %d", resp.StatusCode) // status error
	}

	hasSeen := make(map[string]bool) // initialize map to prevent duplicate links
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err
	}

	doc.Find("#bodyContent a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		// exclude certain links
		if exists && strings.HasPrefix(href, "/wiki/") && !strings.HasPrefix(href, "/wiki/File:") && !hasSeen[href] && !strings.HasPrefix(href, "/wiki/Category:") && !strings.HasPrefix(href, "/wiki/Template:") && !strings.HasPrefix(href, "/wiki/Special:") && !strings.HasPrefix(href, "/wiki/Wikipedia:") && !strings.HasPrefix(href, "/wiki/Help:") && !strings.HasPrefix(href, "/wiki/Portal:") && !strings.HasPrefix(href, "/wiki/Template_talk:") {
			*resultData = append(*resultData, "https://en.wikipedia.org"+href)
			hasSeen[href] = true
		}
	})

	return nil
}

// getting scraping concurrently
func GetScrapeLinksConcurrent(link string) []string {
	mutexCache.Lock() // lock cache
	Unique[link] = true
	links, exist := LinkCache[link]
	mutexCache.Unlock() // unlock cache

	if !exist { // if link is not a key in LinkCache, as in the link has not been scraped before
		links = []string{}
		if err := WebScraping(link, &links); err != nil {
			return nil
		}

		mutexCache.Lock() // lock cache
		LinkCache[link] = links
		mutexCache.Unlock()
	}
	return links
}

// record link scraping in json format
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

// read json file
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

// change URL into a wikipedia link
func wikiUrlToTitle(wikiUrl string) string {
	decoded, err := url.QueryUnescape(strings.TrimPrefix(wikiUrl, "https://en.wikipedia.org/wiki/"))
	if err != nil {
		return ""
	}
	return strings.ReplaceAll(decoded, "_", " ")
}

// change path URL to path Title
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

	defer response.Body.Close()

	// Update the URL to the final redirected
	*url = response.Request.URL.String()

	// return bool if status code is 2xx, meaning it has succeded
	return response.StatusCode >= 200 && response.StatusCode < 300
}

// load cache
func LoadCache(filename string) {
	Unique = make(map[string]bool)
	ReadJSON(filename)
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
