package rssgen

import (
	"encoding/xml"
	"log"
	"time"

	"github.com/brianvoe/gofakeit/v5"
)

// Generate a rss-like file and return the content as byte
func Generate(seed int64, size int64) (rssContent []byte) {
	gofakeit.Seed(seed)
	var rss = RSS{
		Channel: ChannelStruct{
			Item: make([]Item, size),
		},
	}
	gofakeit.Struct(&rss)
	rssContent, err := xml.Marshal(&rss)
	if err != nil {
		log.Println(err)
	}
	return
}

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel ChannelStruct
}

type Item struct {
	Title       string    `fake:"{sentence:3}"`
	PubDate     time.Time `fake:"{date}"`
	Link        string    `fake:"{url}"`
	Category    string    `fake:"{bs}"`
	Description string    `fake:"{loremipsumparagraph}"`
}

type ChannelStruct struct {
	Title         string    `fake:"{bs}"`
	Description   string    `fake:"{sentence:8}"`
	Link          string    `fake:"{url}"`
	PubDate       time.Time `fake:"{date}"`
	LastBuildDate time.Time `fake:"{date}"`
	Item          []Item
}
