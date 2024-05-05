package api

import (
	"encoding/json"
	"fmt"
	dbPackage "github.com/bendtheji/go_url_shortener/db"
	"github.com/gorilla/mux"
	"hash/crc32"
	"net/http"
	"os"
)

type CreateShortUrlRequest struct {
	Url         string `json:"long_url"`
	Description string `json:"description"`
}

type GetShortUrlRequest struct {
	Url string `json:"long_url"`
}

func CreateShortUrlHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("ORIGIN_ALLOWED"))

	db, err := dbPackage.ConnectToDB(*dbPackage.DbConfig)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var req CreateShortUrlRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		// TODO: handle decode request error
	}

	// here we check if longURL is already in DB, req.Url

	// hash the longurl
	crc32hashString := fmt.Sprintf("%08x", crc32.Checksum([]byte(req.Url), crc32.IEEETable))

	// check if hash has been used before

	// create row in table with the created hash
	err = dbPackage.CreateShortUrl(r.Context(), db, req.Url, crc32hashString, req.Description)
	if err != nil {
		// TODO: handle create error
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Short URL created")
}

func ListShortUrlHandler(w http.ResponseWriter, r *http.Request) {
	db, err := dbPackage.ConnectToDB(*dbPackage.DbConfig)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	list, err := dbPackage.ListShortUrls(r.Context(), db)
	if err != nil {
		// TODO: handle this
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("ORIGIN_ALLOWED"))
	json.NewEncoder(w).Encode(list)
	w.WriteHeader(http.StatusOK)
}

func GetShortUrlHandler(w http.ResponseWriter, r *http.Request) {
	db, err := dbPackage.ConnectToDB(*dbPackage.DbConfig)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Get the 'id' parameter from the URL
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Call the GetUser function to fetch the user data from the database
	longUrl, err := dbPackage.GetShortUrl(r.Context(), db, idStr)
	if err != nil {
		// TODO: handle this
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("ORIGIN_ALLOWED"))
	http.Redirect(w, r, longUrl, 301)
}
