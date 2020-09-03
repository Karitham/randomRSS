package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Karitham/randomRSS/rssgen"
)

func main() {
	const name = "random.rss"
	log.SetFlags(log.Lshortfile)

	// Useless
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "to get a rss feed go to http://localhost:8080/rss")
	})

	// Return a fixed content
	http.HandleFunc("/rss", func(w http.ResponseWriter, req *http.Request) {
		modtime := time.Now()
		content := RSS(1)
		http.ServeContent(w, req, name, modtime, content)
	})

	// Return a random content
	http.HandleFunc("/fuzz", func(w http.ResponseWriter, req *http.Request) {
		seed, err := parseRequestURI(req.RequestURI)
		if err != nil {
			fmt.Fprintf(w, "Wrong URL format, try url:port/fuzz?seed=seed")
		}
		modtime := time.Now()
		content := RSS(int64(seed))
		http.ServeContent(w, req, name, modtime, content)
	})

	log.Println("Listening on http://localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// RSS generate the file and return return it
func RSS(seed int64) (xml io.ReadSeeker) {
	var rss rssgen.RSS
	content := rssgen.Generate(&rss, seed)
	return bytes.NewReader(content)
}

func parseRequestURI(uri string) (seed int, err error) {
	seed, err = strconv.Atoi((strings.TrimPrefix(uri, "/fuzz?seed=")))
	if seed == 0 {
		return 0, errors.New("invalid seed")
	}
	return
}
