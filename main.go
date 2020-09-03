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
	Launch()
}

// Launch runs the server
func Launch() {
	log.SetFlags(log.Lshortfile)
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "to get a rss feed to to http://localhost:8080/rss")
	})
	http.HandleFunc("/rss", func(w http.ResponseWriter, req *http.Request) {
		// tell the browser the returned content should be downloaded
		modtime := time.Now()
		content := RSS()
		const name = "rss"
		w.Header().Add("Content-Disposition", "Attachment")
		http.ServeContent(w, req, name, modtime, content)
	})
	log.Println("Listening on http://localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// RSS is used to create the file and return it as byte
func RSS() io.ReadSeeker {
	var rss rssgen.Feed
	content := rssgen.Generate(&rss)
	return bytes.NewReader(content)
}
