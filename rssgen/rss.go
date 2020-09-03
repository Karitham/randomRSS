package rssgen

import (
	"encoding/xml"
	"log"

	"github.com/brianvoe/gofakeit"
)

// Generate creates the RSS file and dumps it into the specified directory
func Generate(xmlStruct *Feed) (rssContent []byte) {
	gofakeit.Seed(0)
	gofakeit.Struct(&xmlStruct)
	rssContent, err := xml.Marshal(&xmlStruct)
	if err != nil {
		log.Println(err)
	}
	return
}

// Feed represent a random XML feed
type Feed struct {
	XMLName  string `fake:"{name}"`
	Text     string `fake:"{sentence:25}"`
	Media    string `fake:"{fruit}"`
	Category struct {
		Text  string `fake:"{quote}"`
		Term  string `fake:"{color}"`
		Label string `fake:"{vehicle}"`
	} `fake:"{snack}"`
	Updated string `fake:"{bool}"`
	Icon    string `fake:"{url}"`
	ID      string `fake:"{uuid}"`
	Link    [25]struct {
		Text string `fake:"{quote}"`
		Rel  string `fake:"{state}"`
		Href string `fake:"{timezone}"`
		Type string `fake:"{beerstyle}"`
	} `fake:"{url}"`
	Subtitle string `fake:"{quote}"`
	Title    string `fake:"{quote}"`
	Entry    [25]struct {
		Text   string `fake:"{sentence:15}"`
		Author struct {
			Text string `fake:"{sentence:15}"`
			Name string `fake:"{name}"`
			URI  string `fake:"{url}"`
		} `fake:"{username}"`
		Category struct {
			Text  string `fake:"{quote}"`
			Term  string `fake:"{color}"`
			Label string `fake:"{beer}"`
		} `fake:"{snack}"`
		Content struct {
			Text string `fake:"{sentence:15}"`
			Type string
		} `fake:"{diner}"`
		ID        string `fake:"{uuid}"`
		Thumbnail struct {
			Text string `fake:"{sentence:15}"`
			URL  string `fake:"{imageurl}"`
		} `fake:"{url}"`
		Link    string `fake:"{url}"`
		Updated string `fake:"{bool}"`
		Title   string `fake:"{bool}"`
	} `fake:"{fruit}"`
}
