package main

import (
	"database/sql"
	"log"
	"net/http"
)

type URLData struct {
	OriginalURL string `json:"original_url"`
	ShortURL    string `json:"short_url"`
	ClickCount  int    `json:"click_count"`
}

var db *sql.DB

func main() {
	// API endpoints
	http.HandleFunc("/api/shorten", shortenURLHandler) // create a short URL
	http.HandleFunc("/s/", redirectHandler)            // redirect to the original URL
	http.HandleFunc("/api/stats", statsHandler)        // get URL stats

	log.Printf("Server is running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func shortenURLHandler(w http.ResponseWriter, r *http.Request) {

}

func redirectHandler(w http.ResponseWriter, r *http.Request) {

}

func statsHandler(w http.ResponseWriter, r *http.Request) {

}
