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
	timeFmt := time.RFC822Z
	var rss = &RSS{
		Version: "2.0",
		Channel: ChannelStruct{
			Item:          make([]Item, size),
			LastBuildDate: gofakeit.Date().Format(timeFmt),
			PubDate:       gofakeit.Date().Format(timeFmt),
		},
	}
	gofakeit.Struct(&rss)
	for i := range rss.Channel.Item {
		rss.Channel.Item[i].PubDate = gofakeit.Date().Format(timeFmt)
	}
	rssContent, err := xml.Marshal(&rss)
	if err != nil {
		log.Println(err)
	}
	return
}

type RSS struct {
	XMLName xml.Name      `xml:"rss"`
	Version string        `fake:"skip" xml:"version,attr"`
	Channel ChannelStruct `xml:"channel"`
}

type Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title" fake:{sentence:3}`
	PubDate     string   `fake:"skip" xml:"pubDate"`
	Link        string   `fake:"{url}" xml:"link"`
	Category    string   `fake:"{bs}" xml:"category"`
	Description string   `fake:"{loremipsumparagraph}" xml:"description"`
}

type ChannelStruct struct {
	Title         string `fake:"{bs}" xml:"title"`
	Description   string `fake:"{sentence:8}" xml:"description"`
	Link          string `fake:"{url}" xml:"link"`
	PubDate       string `fake:"skip" xml:"pubDate"`
	LastBuildDate string `fake:"skip" xml:"lastBuildDate"`
	Item          []Item
}
