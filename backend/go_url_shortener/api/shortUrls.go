package api

import (
	"encoding/json"
	"fmt"
	dbPackage "github.com/bendtheji/go_url_shortener/db"
	apiError "github.com/bendtheji/go_url_shortener/errors"
	"github.com/gorilla/mux"
	"hash/crc32"
	"net/http"
	"net/url"
	"os"
)

type CreateShortUrlRequest struct {
	Url         string `json:"long_url"`
	Description string `json:"description"`
}

type CreateShortUrlResponse struct {
	ShortUrl string `json:"short_url"`
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
		apiError.HandleApiError(w, apiError.HandleError(fmt.Errorf("%w: %w", apiError.ReqUnmarshalTypeErr, err)))
		return
	}

	if req.Url == "" {
		apiError.HandleApiError(w, apiError.HandleError(apiError.MissingLongURL))
		return
	}

	parsedUrl, err := url.Parse(req.Url)
	if err != nil || !parsedUrl.IsAbs() {
		apiError.HandleApiError(w, apiError.HandleError(apiError.InvalidLongURL))
		return
	}

	if req.Description == "" {
		apiError.HandleApiError(w, apiError.HandleError(apiError.MissingDescription))
		return
	}

	// check if long url exists
	found, err := dbPackage.GetShortUrlByLongUrl(r.Context(), db, req.Url)
	if found != "" {
		apiError.HandleApiError(w, apiError.HandleError(apiError.DuplicateLongURL))
		return
	}

	longUrl := req.Url
	// hash the longurl
	var crc32hashString string
	for {
		crc32hashString = fmt.Sprintf("%08x", crc32.Checksum([]byte(longUrl), crc32.IEEETable))

		// check if hash has been used before
		// if hasn't been used before, then we break out of the loop
		// else we append the hash string to the long url and hash it again
		found, err = dbPackage.GetShortUrl(r.Context(), db, crc32hashString)
		if found == "" {
			break
		}
		longUrl = longUrl + ":" + crc32hashString
	}

	// create row in table with the created hash
	err = dbPackage.CreateShortUrl(r.Context(), db, req.Url, crc32hashString, req.Description)
	if err != nil {
		apiError.HandleApiError(w, err)
		return
	}

	res := CreateShortUrlResponse{ShortUrl: crc32hashString}
	jsonData, err := json.Marshal(res)
	if err != nil {
		apiError.HandleApiError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(jsonData)
}

func ListShortUrlHandler(w http.ResponseWriter, r *http.Request) {
	db, err := dbPackage.ConnectToDB(*dbPackage.DbConfig)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	list, err := dbPackage.ListShortUrls(r.Context(), db)
	if err != nil {
		apiError.HandleApiError(w, err)
		return
	}

	jsonData, err := json.Marshal(list)
	if err != nil {
		apiError.HandleApiError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("ORIGIN_ALLOWED"))
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func GetShortUrlHandler(w http.ResponseWriter, r *http.Request) {
	db, err := dbPackage.ConnectToDB(*dbPackage.DbConfig)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Get the 'shortUrlHash' parameter from the URL
	vars := mux.Vars(r)
	shortUrlHash := vars["shortUrlHash"]

	// Call the GetUser function to fetch the user data from the database
	longUrl, err := dbPackage.GetShortUrl(r.Context(), db, shortUrlHash)
	if err != nil {
		apiError.HandleApiError(w, err)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("ORIGIN_ALLOWED"))
	http.Redirect(w, r, longUrl, 301)
}
