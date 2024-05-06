package api

import (
	"encoding/json"
	"fmt"
	dbPackage "github.com/bendtheji/go_url_shortener/db"
	apiError "github.com/bendtheji/go_url_shortener/errors"
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
		apiError.HandleApiError(w, apiError.HandleError(fmt.Errorf("%w: %w", apiError.ReqUnmarshalTypeErr, err)))
		return
	}

	if req.Url == "" {
		apiError.HandleApiError(w, apiError.HandleError(apiError.MissingLongURL))
		return
	}

	if req.Description == "" {
		apiError.HandleApiError(w, apiError.HandleError(apiError.MissingDescription))
		return
	}

	// check if long url exists
	found, err := dbPackage.GetLongUrl(r.Context(), db, req.Url)
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
		apiError.HandleApiError(w, err)
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
		apiError.HandleApiError(w, err)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("ORIGIN_ALLOWED"))
	http.Redirect(w, r, longUrl, 301)
}
