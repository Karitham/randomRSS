package rssgen

import (
	"encoding/xml"
	"log"
	"time"

	"github.com/brianvoe/gofakeit/v5"
)

// Generate a rss-like file and return the content as byte
func Generate(xmlStruct *RSS, seed int64) (rssContent []byte) {
	gofakeit.Seed(seed)
	gofakeit.Struct(&xmlStruct)
	rssContent, err := xml.Marshal(&xmlStruct)
	if err != nil {
		log.Println(err)
	}
	// return append([]byte(`<rss xmlns:atom="http://www.w3.org/2005/Atom" version="2.0">`), rssContent...)
	return
}

// Feed represent a random RSS file, this is not perfect and may be improved.
// For now it's just a random jumble, I may later improve this to ressemble an actual feed

type Item struct {
	Title       string    `fake:"{sentence:3}"`
	PubDate     time.Time `fake:"{date}"`
	Link        string    `fake:"{url}"`
	Category    string    `fake:"{bs}"`
	Description string    `fake:"{loremipsumparagraph}"`
}

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel struct {
		Title         string    `fake:"{bs}"`
		Description   string    `fake:"{sentence:8}"`
		Link          string    `fake:"{url}"`
		PubDate       time.Time `fake:"{date}"`
		LastBuildDate time.Time `fake:"{date}"`
		Item          []Item
	}
}
