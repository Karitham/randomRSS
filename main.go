package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Karitham/randomRSS/rssgen"
)

func main() {
	log.SetFlags(log.Lshortfile)

	// Useless
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "to get a rss feed go to http://localhost:8080/rss")
	})

	// Return a random content
	http.HandleFunc("/rss", func(w http.ResponseWriter, req *http.Request) {
		var seed int64 = 1
		var size int64 = 20
		var err error

		// Get query parameters
		seedStr := req.FormValue("seed")
		sizeStr := req.FormValue("size")

		if seedStr != "" {
			seed, err = strconv.ParseInt(seedStr, 10, 0)
			if err != nil {
				w.WriteHeader(400)
				fmt.Fprintf(w, "%s", err)
				return
			}
		}

		if sizeStr != "" {
			size, err = strconv.ParseInt(sizeStr, 10, 0)
			if err != nil {
				w.WriteHeader(400)
				fmt.Fprintf(w, "%s", err)
				return
			}
		}

		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", rssgen.Generate(seed, size))
	})

	log.Println("Listening on http://localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
