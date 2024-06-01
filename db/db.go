package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type Url struct {
	ID          int64
	OriginalURL string
	ShorterURL  string
	CreatedAt   time.Time
}

func ConnectDB() (*sql.DB, error) {
	connStr := "user=akmal password=12345678 dbname=ghortener sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func GetOriginalUrlByShorter(url string) (Url, error) {
	var urlData Url
	db, err := ConnectDB()
	if err != nil {
		fmt.Printf("Error when trying to connect to DB: %s", err)
		return urlData, err
	}

	row := db.QueryRow(`SELECT id, original_url, shorter_url FROM urls WHERE shorter_url = $1`, url)

	err = row.Scan(&urlData.ID, &urlData.OriginalURL, &urlData.ShorterURL)
	if err != nil {
		return urlData, err
	}

	return urlData, nil
}

func CreateShorterUrl(url Url) (string, error) {
	var newId int64
	db, err := ConnectDB()
	if err != nil {
		fmt.Printf("Error when trying to connect to DB: %s", err)
		return "", err
	}

	err = db.QueryRow("INSERT INTO urls (original_url, shorter_url, created_at) VALUES ($1, $2, $3) RETURNING id", url.OriginalURL, url.ShorterURL, url.CreatedAt).Scan(&newId)
	if err != nil {
		return "", err
	}

	return url.ShorterURL, nil
}
