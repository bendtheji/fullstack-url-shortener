package db

import (
	"context"
	"database/sql"
	"time"
)

func CreateShortUrl(ctx context.Context, db *sql.DB, url string, hashString string, description string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := "INSERT INTO urls (short_url, long_url, description) VALUES (?, ?, ?)"
	_, err := db.ExecContext(ctx, query, hashString, url, description)
	if err != nil {
	}
	return nil
}

type Url struct {
	ID          int
	ShortUrl    string
	LongUrl     string
	Description string
}

func GetShortUrl(ctx context.Context, db *sql.DB, str string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := "SELECT * FROM urls WHERE short_url = ?"
	row := db.QueryRowContext(ctx, query, str)

	url := &Url{}
	err := row.Scan(&url.ID, &url.ShortUrl, &url.LongUrl, &url.Description)
	if err != nil {
	}
	return url.LongUrl, nil
}

func ListShortUrls(ctx context.Context, db *sql.DB) ([]Url, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := "SELECT * FROM urls"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
	}

	urls := make([]Url, 0)
	for rows.Next() {
		url := &Url{}
		err := rows.Scan(&url.ID, &url.ShortUrl, &url.LongUrl, &url.Description)
		if err != nil {
		}
		urls = append(urls, *url)
	}
	return urls, nil
}
