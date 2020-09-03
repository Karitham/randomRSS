package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/Karitham/randomRSS/rssgen"
)

func main() {
	const name = "rss"
	log.SetFlags(log.Lshortfile)

	// Useless
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "to get a rss feed go to http://localhost:8080/rss")
	})

	// Send a downloadable content
	http.HandleFunc("/rss", func(w http.ResponseWriter, req *http.Request) {
		modtime := time.Now()
		content := RSS()
		// tell the browser the returned content should be downloaded
		w.Header().Add("Content-Disposition", "Attachment")
		http.ServeContent(w, req, name, modtime, content)
	})

	log.Println("Listening on http://localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// RSS generate the file and return return it
func RSS() (xml io.ReadSeeker) {
	var rss rssgen.Feed
	content := rssgen.Generate(&rss)
	return bytes.NewReader(content)
}
