package api

import (
	"bytes"
	"fmt"
	"github.com/bendtheji/go_url_shortener/db"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func setDBEnvConfig(t *testing.T) {
	t.Setenv("DB_HOST", "localhost")
	t.Setenv("DB_PORT", "3306")
	t.Setenv("DB_USERNAME", "mysql")
	t.Setenv("DB_PASSWORD", "password")
	t.Setenv("DB_DATABASE", "sample_url_shortener")
}

func setWrongDBEnvConfig(t *testing.T) {
	t.Setenv("DB_HOST", "localhost")
	t.Setenv("DB_PORT", "3306")
	t.Setenv("DB_USERNAME", "mysql")
	t.Setenv("DB_PASSWORD", "password")
	t.Setenv("DB_DATABASE", "wrong_db_name")
}

func clearTable() {
	db, _ := db.ConnectToDB(*db.DbConfig)

	db.Exec("DELETE FROM urls")
	db.Exec("ALTER TABLE urls AUTO_INCREMENT = 1")
}

func TestCreateShortUrlHandler_happy_path(t *testing.T) {
	setDBEnvConfig(t)
	db.InitDbConfig()

	clearTable()

	payload := []byte(`{"long_url":"https://www.reddit.com/r/drums/","description":"Drums subreddit"}`)
	req, err := http.NewRequest("POST", "/urls", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateShortUrlHandler)

	handler.ServeHTTP(rr, req)

	fmt.Printf("%v", rr)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "Short URL created"
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	clearTable()
}

func TestCreateShortUrlHandler_missing_long_url(t *testing.T) {
	setDBEnvConfig(t)
	db.InitDbConfig()

	clearTable()

	payload := []byte(`{"description":"Drums subreddit"}`)
	req, err := http.NewRequest("POST", "/urls", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateShortUrlHandler)

	handler.ServeHTTP(rr, req)

	fmt.Printf("%v", rr)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	expected := "missing long url"
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	clearTable()
}

func TestCreateShortUrlHandler_missing_description(t *testing.T) {
	setDBEnvConfig(t)
	db.InitDbConfig()

	clearTable()

	payload := []byte(`{"long_url":"https://www.reddit.com/r/drums/"}`)
	req, err := http.NewRequest("POST", "/urls", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateShortUrlHandler)

	handler.ServeHTTP(rr, req)

	fmt.Printf("%v", rr)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	expected := "missing description"
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	clearTable()
}

func TestCreateShortUrlHandler_duplicate_long_url(t *testing.T) {
	setDBEnvConfig(t)
	db.InitDbConfig()

	clearTable()

	handler := http.HandlerFunc(CreateShortUrlHandler)

	// first request
	payload := []byte(`{"long_url":"https://www.reddit.com/r/drums/","description":"Drums subreddit"}`)
	req, err := http.NewRequest("POST", "/urls", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	// duplicate request
	req, err = http.NewRequest("POST", "/urls", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	fmt.Printf("%v", rr)
	if status := rr.Code; status != http.StatusConflict {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusConflict)
	}

	expected := "duplicate long url"
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	clearTable()
}

func TestCreateShortUrlHandler_decode_error(t *testing.T) {
	setDBEnvConfig(t)
	db.InitDbConfig()

	clearTable()

	payload := []byte(`{"long_url":"https://www.reddit.com/r/drums/","description"}`)
	req, err := http.NewRequest("POST", "/urls", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateShortUrlHandler)

	handler.ServeHTTP(rr, req)

	fmt.Printf("%v", rr)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	expected := "invalid request type: invalid character '}' after object key"
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	clearTable()
}

func TestCreateShortUrlHandler_internal_server_error(t *testing.T) {
	setWrongDBEnvConfig(t)
	db.InitDbConfig()

	payload := []byte(`{"long_url":"https://www.reddit.com/r/drums/","description":"Drums subreddit"}`)
	req, err := http.NewRequest("POST", "/urls", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateShortUrlHandler)

	handler.ServeHTTP(rr, req)

	fmt.Printf("%v", rr)
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}
