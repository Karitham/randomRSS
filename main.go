package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"

	"github.com/brianvoe/gofakeit/v5"
)

func main() {
	var feed Feed
	createRRS(&feed, "rss")
}

// createRRS creates the RSS file and dumps it into the same directory
func createRRS(xmlStruct *Feed, filename string) {
	gofakeit.Struct(&xmlStruct)
	f, err := xml.Marshal(&xmlStruct)
	if err != nil {
		log.Println(err)
	}
	ioutil.WriteFile(filename, f, os.ModePerm)
}

// Feed represent a random XML feed
type Feed struct {
	XMLName  string `fake:"{name}"`
	Text     string `fake:"{sentence:5}"`
	Media    string `fake:"{fruit}"`
	Category struct {
		Text  string `fake:"{quote}"`
		Term  string `fake:"{color}"`
		Label string `fake:"{vehicle}"`
	} `fake:"{snack}"`
	Updated string `fake:"{bool}"`
	Icon    string `fake:"{url}"`
	ID      string `fake:"{uuid}"`
	Link    []struct {
		Text string `fake:"{quote}"`
		Rel  string `fake:"{state}"`
		Href string `fake:"{timezone}"`
		Type string `fake:"{beerstyle}"`
	} `fake:"{url}"`
	Subtitle string `fake:"{quote}"`
	Title    string `fake:"{quote}"`
	Entry    [15]struct {
		Text   string `fake:"{sentence:5}"`
		Author struct {
			Text string `fake:"{sentence:5}"`
			Name string `fake:"{name}"`
			URI  string `fake:"{url}"`
		} `fake:"{username}"`
		Category struct {
			Text  string `fake:"{quote}"`
			Term  string `fake:"{color}"`
			Label string `fake:"{vehicle}"`
		} `fake:"{snack}"`
		Content struct {
			Text string `fake:"{sentence:5}"`
			Type string
		} `fake:"{diner}"`
		ID        string `fake:"{uuid}"`
		Thumbnail struct {
			Text string `fake:"{sentence:5}"`
			URL  string `fake:"{imageurl}"`
		} `fake:"{url}"`
		Link    string `fake:"{url}"`
		Updated string `fake:"{bool}"`
		Title   string `fake:"{bool}"`
	} `fake:"{fruit}"`
}
