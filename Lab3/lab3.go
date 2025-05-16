package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Feed struct {
	URL      string
	LastRead time.Time
	Articles int // Simulating the number of articles fetched
}

func fetchFeedData(feedURL string) Feed {
	rand.Seed(time.Now().UnixNano())
	return Feed{
		URL:      feedURL,
		LastRead: time.Now(),
		Articles: rand.Intn(20),
	}
}

func main() {

	var rssFeeds [3]string = [3]string{
		"https://example.com/feed1.xml",
		"https://example.com/feed2.xml",
		"https://example.com/feed3.xml",
	}

	var fetchedFeeds []Feed

	for _, url := range rssFeeds {
		fetchedFeeds = append(fetchedFeeds, fetchFeedData(url))
	}

	fmt.Println("Fetched RSS Feeds Metadata:")
	for _, feed := range fetchedFeeds {
		fmt.Printf("- URL: %s\n  Last Read: %v\n  Articles Fetched: %d\n\n",
			feed.URL, feed.LastRead.Format(time.RFC1123), feed.Articles)
	}
}
