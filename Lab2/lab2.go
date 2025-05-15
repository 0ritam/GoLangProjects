package main

import (
	"fmt"
	"net/http"
	"strings"
)

func isValidURL(url string) bool {
	return strings.HasPrefix(url, "http")
}

func checkRSSFeed(url string) {
	if !isValidURL(url) {
		fmt.Println(url, "❌ Invalid URL format")
		return
	}

	// Send a request to check if the URL is accessible
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		fmt.Println(url, "❌ Unreachable RSS Feed")
		return
	}

	fmt.Println(url, "✅ Valid RSS Feed")
}

func main() {

	urls := []string{
		"https://rss.nytimes.com/services/xml/rss/nyt/World.xml", // Valid
		"htp://invalid-url.com",                                  // Invalid format
		"https://example.com/notrss",                             // Might not be a real feed
	}

	for _, url := range urls {
		checkRSSFeed(url)
	}
}
