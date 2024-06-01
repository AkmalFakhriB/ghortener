package main

import (
	"log"
	"net/http"

	API "github.com/AkmalFakhriB/ghortener/api"
)

func main() {
	http.HandleFunc("/{url}", API.RedirectToOriginalUrl)
	http.HandleFunc("/newurl", API.CreateShorterUrl)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
