package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"
)

type FeedFetcher interface {
	FetchData() (string, error)
	GetType() string
}

// XMLFeed struct
type XMLFeed struct {
	URL string
}

// JSONFeed struct
type JSONFeed struct {
	URL string
}

// AtomFeed struct
type AtomFeed struct {
	URL string
}

type RSSItem struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
}

type RSSFeed struct {
	Items []RSSItem `xml:"channel>item"`
}

type AtomEntry struct {
	Title   string `xml:"title"`
	Summary string `xml:"summary"`
	Link    struct {
		Href string `xml:"href,attr"`
	} `xml:"link"`
	Published string `xml:"published"`
}

type AtomFeedStruct struct {
	Entries []AtomEntry `xml:"entry"`
}

// XMLFeed implementation
func (x XMLFeed) FetchData() (string, error) {
	start := time.Now()
	defer func() {
		fmt.Printf("Time taken to fetch XML data: %v\n", time.Since(start))
	}()

	resp, err := http.Get(x.URL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var rssFeed RSSFeed
	err = xml.Unmarshal(body, &rssFeed)
	if err != nil {
		return "", err
	}

	items := rssFeed.Items
	if len(items) > 3 {
		items = items[:3]
	}

	return fmt.Sprintf("Fetched XML data: %+v", items), nil
}

func (x XMLFeed) GetType() string {
	return "XML"
}

// JSONFeed implementation
func (j JSONFeed) FetchData() (string, error) {
	start := time.Now()
	defer func() {
		fmt.Printf("Time taken to fetch JSON data: %v\n", time.Since(start))
	}()

	resp, err := http.Get(j.URL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result []map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	if len(result) > 3 {
		result = result[:3]
	}

	return fmt.Sprintf("Fetched JSON data: %+v", result), nil
}

func (j JSONFeed) GetType() string {
	return "JSON"
}

// AtomFeed implementation
func (a AtomFeed) FetchData() (string, error) {
	start := time.Now()
	defer func() {
		fmt.Printf("Time taken to fetch Atom data: %v\n", time.Since(start))
	}()

	resp, err := http.Get(a.URL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var atomFeed AtomFeedStruct
	err = xml.Unmarshal(body, &atomFeed)
	if err != nil {
		return "", err
	}

	entries := atomFeed.Entries
	if len(entries) > 3 {
		entries = entries[:3]
	}

	return fmt.Sprintf("Fetched Atom data: %+v", entries), nil
}

func (a AtomFeed) GetType() string {
	return "Atom"
}

// demonstrating polymorphism
func FetchAndPrint(feed FeedFetcher) {
	feedType := feed.GetType()
	fmt.Printf("Fetching %s feed...\n", feedType)

	data, err := feed.FetchData()
	if err != nil {
		fmt.Printf("Error fetching data: %v\n", err)
		return
	}
	fmt.Println(data)
}

func main() {

	xmlFeed := XMLFeed{URL: "https://feeds.bbci.co.uk/news/rss.xml"}
	jsonFeed := JSONFeed{URL: "https://api.github.com/users/octocat/repos"}
	atomFeed := AtomFeed{URL: "https://hnrss.org/frontpage?format=atom"}

	FetchAndPrint(xmlFeed)
	FetchAndPrint(jsonFeed)
	FetchAndPrint(atomFeed)
}
