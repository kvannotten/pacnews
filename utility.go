package main

import (
	"fmt"
	"os"

	"github.com/boltdb/bolt"
	"github.com/fatih/color"
	"github.com/jaytaylor/html2text"
	"github.com/mmcdole/gofeed"
	"github.com/urfave/cli"
)

func feedItems() []*gofeed.Item {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://www.archlinux.org/feeds/news/")
	if err != nil {
		return []*gofeed.Item{}
	}

	for i, j := 0, len(feed.Items)-1; i < j; i, j = i+1, j-1 {
		feed.Items[i], feed.Items[j] = feed.Items[j], feed.Items[i]
	}

	return feed.Items
}

func checkNews(c *cli.Context) error {
	new := false

	DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("news"))
		for _, item := range feedItems() {
			v := b.Get([]byte(item.GUID))
			if v == nil {
				new = true
			}
		}
		return nil
	})

	if new {
		color.New(color.FgRed).Println("\n\n>>>>>>>>>> There are new news items. Read them before you continue <<<<<<<<<<\n\n")
		os.Exit(1)
	}
	return nil
}

func readNews(c *cli.Context) error {
	DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("news"))
		for _, item := range feedItems() {
			if b.Get([]byte(item.GUID)) == nil {
				print(item)
				b.Put([]byte(item.GUID), []byte(item.Title))
			}
		}
		return nil
	})

	return nil
}

func print(news *gofeed.Item) {
	color.New(color.FgRed, color.Bold).Printf("\n%s\n", news.Title)
	color.New(color.FgBlue).Printf("%s\n", news.Link)
	color.New(color.FgCyan).Printf("%s\n", news.Updated)
	text, _ := html2text.FromString(news.Description)
	color.White(fmt.Sprintf("\t%s\n", text))
}
