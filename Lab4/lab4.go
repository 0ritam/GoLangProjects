package main

import (
	"fmt"
)

type RSSFeed struct {
	Title       string
	URL         string
	Description string
}

type User struct {
	Name       string
	Subscribed map[string][]RSSFeed // Category -> List of feeds
}

func (u *User) Subscribe(category string, feed RSSFeed) {
	if u.Subscribed == nil {
		u.Subscribed = make(map[string][]RSSFeed)
	}
	u.Subscribed[category] = append(u.Subscribed[category], feed)
}

func (u *User) DisplaySubscriptions() {
	fmt.Printf("\n%s's Subscribed Feeds:\n", u.Name)
	for category, feeds := range u.Subscribed {
		fmt.Println("Category:", category)
		for _, feed := range feeds {
			fmt.Printf(" - %s (%s): %s\n", feed.Title, feed.URL, feed.Description)
		}
	}
}

func main() {

	user := User{Name: "Alice"}

	user.Subscribe("Technology", RSSFeed{"TechCrunch", "https://techcrunch.com/feed/", "Latest technology news"})
	user.Subscribe("Technology", RSSFeed{"Hacker News", "https://news.ycombinator.com/rss", "Tech discussions and startups"})
	user.Subscribe("Sports", RSSFeed{"ESPN", "https://www.espn.com/espn/rss/news", "Latest sports updates"})

	user.DisplaySubscriptions()
}
