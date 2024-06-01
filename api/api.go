package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	DB "github.com/AkmalFakhriB/ghortener/db"
	helper "github.com/AkmalFakhriB/ghortener/helper"
)

func RedirectToOriginalUrl(w http.ResponseWriter, r *http.Request) {
	url := r.PathValue("url")

	urlData, err := DB.GetOriginalUrlByShorter(url)
	if err != nil {
		fmt.Printf("Error when trying to get original url data: %s", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Add page loading before redirecting
	http.Redirect(w, r, urlData.OriginalURL, http.StatusSeeOther)
}

func CreateShorterUrl(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error parsing form data: %v", err)
		return
	}

	shortRandom := r.FormValue("shorterUrl")
	if r.FormValue("shorterUrl") == "" {
		shortRandom = helper.RandomString(5)
	}

	dateCreated := time.Now()

	url := DB.Url{
		OriginalURL: r.FormValue("originalUrl"),
		ShorterURL:  shortRandom,
		CreatedAt:   dateCreated,
	}

	shorterUrl, err := DB.CreateShorterUrl(url)
	if err != nil {
		fmt.Printf("Error while calling db CreateShorterUrl: %s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(shorterUrl)
}
