package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

// Structs to parse RSS Feed
type RSS struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Title string `xml:"title"`
	Items []Item `xml:"item"`
}

type Item struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
}

// Function to fetch and parse RSS Feed
func fetchRSSFeed(url string) (*RSS, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch RSS feed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response status: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var rss RSS
	err = xml.Unmarshal(data, &rss)
	if err != nil {
		return nil, fmt.Errorf("failed to parse RSS feed: %v", err)
	}

	return &rss, nil
}

func main() {

	url := "https://news.ycombinator.com/rss"

	rss, err := fetchRSSFeed(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Feed Title:", rss.Channel.Title)
	fmt.Println("\nLatest Articles:")
	for _, item := range rss.Channel.Items {
		fmt.Printf("- %s (%s)\n", item.Title, item.Link)
	}
}
