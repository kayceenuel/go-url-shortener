package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/Home", handleHome)
	http.HandleFunc("/shorten", handleshorten)
	http.HandleFunc("/redirect", handleRedirect)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleshorten(w http.ResponseWriter, r *http.Request) {

}

func handleRedirect(w http.ResponseWriter, r *http.Request) {

}

func handleHome(w http.ResponseWriter, r *http.Request) {

}
