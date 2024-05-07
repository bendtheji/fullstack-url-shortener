# URL Shortener Fullstack Application

## Technologies used
1) Go
2) MySQL
3) gorilla/mux
4) React
5) Tailwind CSS
6) Docker

## Clone Repo
```
git clone https://github.com/bendtheji/fullstack-url-shortener.git
```

## Run with Docker (Recommended)
You can use Docker to start up this application. Navigate to the root project with `docker-compose.yml` then run the command below:
```
docker compose up -d
```
You'll need to give it a few seconds to start up because of the database migration scripts.

Then, on your browser, access `http://localhost:3000` and you can start using the app.


## Run locally
If you want to run it locally, you'll need to setup the things below and follow the corresponding subfolders' `README.md` file to complete the setup.

### Install required tools
Install Go: https://go.dev/doc/install

Install MySQL: https://dev.mysql.com/downloads/installer/

Install Node and npm: https://docs.npmjs.com/downloading-and-installing-node-js-and-npm

### Setup MySQL Database
Enter MySQL shell as root:
```
mysql -u root -p
```
You'll need the password used during MySQL installation to log into the shell.

Create Database:
```
create database sample_url_shortener;
```

Create user for interacting with DB:
```
create user 'mysql'@'localhost' identified by 'password';
```

Grant privileges for user to the DB:
```
grant all privileges on sample_url_shortener.* to 'mysql'@'localhost';
```

### DB Migrations
Install `golang-migrate` executable
```
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Set executable as command:
```
export migrate=$HOME/go/bin/migrate
```

Create directory with SQL script (can skip since folder is created):
```
$migrate create -ext sql -dir migrations -seq create_urls_table_something
```

Set MYSQL URL env var:
```
export MYSQL_URL="mysql://<DB_USERNAME>:<DB_PASSWORD>@tcp(localhost:3306)/<DB_NAME>"
```
Replace the placeholders with actual values.

Apply up migrations:
```
$migrate -database ${MYSQL_URL} -path migrations up
```

Apply down migrations:
```
$migrate -database ${MYSQL_URL} -path migrations down
```

Complete the rest of the setup in the respective subfolders.