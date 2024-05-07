package main

import (
	"github.com/bendtheji/go_url_shortener/api"
	"github.com/bendtheji/go_url_shortener/db"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.InitDbConfig()
	r := mux.NewRouter()

	r.HandleFunc("/shortUrls", api.CreateShortUrlHandler).Methods("POST")
	r.HandleFunc("/shortUrls", api.ListShortUrlHandler).Methods("GET")
	r.HandleFunc("/shortUrls/{shortUrlHash}", api.GetShortUrlHandler).Methods("GET")

	log.Println("Server listening on :8090")
	log.Fatal(http.ListenAndServe(":8090", corsMiddleware(r)))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middleware", r.Method)

		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Origin", os.Getenv("ORIGIN_ALLOWED"))
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token, Authorization")
			w.Header().Set("Content-Type", "application/json")
			return
		}

		next.ServeHTTP(w, r)
		log.Println("Executing middleware again")
	})
}
