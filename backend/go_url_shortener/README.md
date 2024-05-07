# Go Backend Server for URL Shortener

## Install Go and MySQL
Install Go: https://go.dev/doc/install

Install MySQL: https://dev.mysql.com/downloads/installer/

Ensure that the DB setup is done as how the `README.md` file dictated.

## Setup for backend

### Create .env file
```
touch .env
echo DB_HOST=localhost >> .env
echo DB_PORT=3306 >> .env
echo DB_USERNAME=<DB_USERNAME> >> .env
echo DB_PASSWORD=<DB_PASSWORD> >> .env
echo DB_DATABASE=<DB_DATABASE> >> .env
```

### Start the application
```
go run .
```

## APIs
List of available endpoints.

### Postman
Import the "URL Shortener.postman_collection.json" file into Postman to get the list of available endpoints.

### Create Short URL
`POST /shortUrls`

Request Body:
```
{
    "long_url": "https://www.reddit.com/r/drums/",
    "description": "Drums subreddit"
}
```

Response Body:
```
{"short_url":"c7dbb529"}
```

### Get Redirect URL
`GET /shortUrls/{shortUrlHash}`

Path params:
- `shortUrlHash`

Response:

Will be redirected to long URL tied to short URL hash.


### List Short URLs
`GET /shortUrls`

Response:
```
[
    {
        "ID": 1,
        "ShortUrl": "332073bb",
        "LongUrl": "https://www.youtube.com/watch?v=fbPmLhM9EZQ",
        "Description": "Conan makes NYC pizza"
    },
    {
        "ID": 2,
        "ShortUrl": "c7dbb529",
        "LongUrl": "https://www.reddit.com/r/drums/",
        "Description": "Drums subreddit"
    }
]
```

## Testing
Make sure MySQL server is started with respective database and table.

Also, go to `shortUrls_test.go` and setup environment variables for testing under the `setDBEnvConfig` function.

If these are not done the tests will fail.

Run the command below to run tests:
```
go test ./...
```

Tests are written for the API handlers at this point but more tests can be written for the DB and the utility functions as well.