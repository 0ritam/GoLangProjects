package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

// RSS structure for XML parsing
type RSS struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Title string `xml:"title"`
	Items []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
}

// Function demonstrating call by value (does not modify original)
func modifyTitleByValue(feed RSS) {
	feed.Channel.Title = "Modified Title (By Reference)"
}

// Function demonstrating call by reference (modifies original)
func modifyTitleByReference(feed *RSS) {
	feed.Channel.Title = "RSS Feed Aggregator By POINTER*️⃣"
}

// Fetch and parse RSS feed
func fetchRSS(url string) (*RSS, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rss RSS
	err = xml.Unmarshal(body, &rss)
	if err != nil {
		return nil, err
	}

	return &rss, nil
}

var tmpl = template.Must(template.New("rssTemplate").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Channel.Title}}</title>
    <style>
        * {
            box-sizing: border-box;
            margin: 0;
            padding: 0;
            font-family: Arial, sans-serif;
        }
        body {
            background-color: #121212;
            color: #ffffff;
            margin: 20px;
        }
        h1 {
            text-align: center;
            margin-bottom: 20px;
            font-size: 2rem;
        }
        .container {
            max-width: 800px;
            margin: auto;
        }
        .item {
            background: #1e1e1e;
            padding: 15px;
            margin: 10px 0;
            border-radius: 8px;
            box-shadow: 0px 4px 6px rgba(255, 255, 255, 0.1);
            transition: transform 0.2s ease-in-out;
        }
        .item:hover {
            transform: scale(1.02);
        }
        .item h2 {
            font-size: 1.5rem;
            color: #ffcc00;
        }
        .item p {
            margin: 10px 0;
            font-size: 1rem;
            line-height: 1.5;
        }
        .item a {
            text-decoration: none;
            color: #66b3ff;
            font-weight: bold;
        }
        .item a:hover {
            color: #ffcc00;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>{{.Channel.Title}}</h1>
        {{range .Channel.Items}}
            <div class="item">
                <h2>{{.Title}}</h2>
                <p>{{.Description}}</p>
                <a href="{{.Link}}" target="_blank">Read more →</a>
            </div>
        {{end}}
    </div>
</body>
</html>
`))

// Handler for displaying RSS feed in HTML
func rssHandler(w http.ResponseWriter, r *http.Request) {
	rss, err := fetchRSS("https://rss.nytimes.com/services/xml/rss/nyt/World.xml") // Replace with desired feed
	if err != nil {
		http.Error(w, "Failed to fetch RSS", http.StatusInternalServerError)
		return
	}

	// Demonstrate call by value (does not affect original)
	modifyTitleByValue(*rss)

	// Demonstrate call by reference (modifies original)
	modifyTitleByReference(rss)

	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, rss)
}

func main() {
	http.HandleFunc("/", rssHandler)
	fmt.Println("Server running on http://localhost:5000...")
	http.ListenAndServe(":5000", nil)
}
