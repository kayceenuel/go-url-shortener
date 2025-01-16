package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/shorten", handleshorten)
	http.HandleFunc("", handleRedirect)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleshorten(w http.ResponseWriter, r *http.Request) {

}

func handleRedirect(w http.ResponseWriter, r *http.Request) {

}
