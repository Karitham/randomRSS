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
		fmt.Fprintf(w, "to get a rss feed go to http://localhost:8080/fuzz?seed=seed")
	})

	// Return a random content
	http.HandleFunc("/fuzz", func(w http.ResponseWriter, req *http.Request) {
		seed, err := strconv.ParseInt(req.FormValue("seed"), 10, 0)
		if err != nil {
			w.WriteHeader(400)
			fmt.Fprintf(w, "%s", err)
		} else {
			w.WriteHeader(200)
			fmt.Fprintf(w, "%s", RSS(seed))
		}
	})

	log.Println("Listening on http://localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// RSS generate the file and return return it
func RSS(seed int64) (xml []byte) {
	var rss rssgen.Feed
	return rssgen.Generate(&rss, seed)
}
