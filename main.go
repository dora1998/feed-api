package main

import (
	"log"
	"os"
	"time"

	"github.com/naoki-kishi/feeder"
)

func main() {
	feed := fetch()

	rss, err := feed.ToRSS()
	if err != nil {
		log.Fatal(err)
		return
	}
	writeFeed("/var/feed/rss.xml", rss)

	atom, err := feed.ToAtom()
	if err != nil {
		log.Fatal(err)
		return
	}
	writeFeed("/var/feed/atom.xml", atom)

	json, err := feed.ToJSON()
	if err != nil {
		log.Fatal(err)
		return
	}
	writeFeed("/var/feed/json.json", json)
}

func fetch() *feeder.Feed {
	rssFetcher := feeder.NewRSSFetcher("http://dorapocket.starfree.jp/feed/")
	qiitaFetcher := feeder.NewQiitaFetcher("https://qiita.com/api/v2/users/dora1998/items")

	// Fetch data using goroutine.
	items := feeder.Crawl(rssFetcher, qiitaFetcher)

	feed := &feeder.Feed{
		Title:       "My feeds",
		Link:        &feeder.Link{Href: "http://feed-api.minoru.dev/rss"},
		Description: "My feeds.",
		Author: &feeder.Author{
			Name:  "Minoru Takeuchi",
			Email: "me@minoru.dev"},
		Created: time.Now(),
		Items:   *items,
	}

	return feed
}

func writeFeed(path string, body string) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	defer file.Close()
	file.Write(([]byte)(body))

	return nil
}
