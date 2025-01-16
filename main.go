package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/shorten", handleShorten)
	http.HandleFunc("/", handleRedirect)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func handleShorten(w http.ResponseWriter, r *http.Request) {}

func handleRedirect(w http.ResponseWriter, r *http.Request) {}
