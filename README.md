# sho.rt 🔗✂️

A lightweight **URL shortener** built with Go that generates short
links, stores mappings in a database, and prevents duplicate entries for
the same original URL.\
Uses **Gin** for the HTTP API and **GORM** as the ORM layer.

## Features

-   Shortens any valid URL
-   Returns existing short URL if already created
-   Persists data in a relational database
-   RESTful API interface
-   Simple, fast, and lightweight

## Technologies

-   Go 1.20+
-   Gin Web Framework
-   GORM (ORM)
-   PostgreSQL (or SQLite/MySQL)
-   godotenv for configuration

## Getting Started

### Prerequisites

-   Go 1.20+
-   PostgreSQL (recommended) or SQLite/MySQL
-   Git

### Setup

1.  Clone the repo:

``` bash
git clone https://github.com/yourusername/short-app.git
cd short-app
```

2.  Create a `.env` file in the project root:

``` txt
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=postgres
DB_NAME=short_db
BASE_URL=http://sho.rt/
```

3.  Install dependencies:

``` bash
go mod tidy
```

4.  Run the app:

``` bash
go run main.go
```

The app runs at:

    http://localhost:8080

------------------------------------------------------------------------

## API Usage

### **Create (or retrieve) a short URL**

**Endpoint:**

    POST /shorten

**Request Body:**

``` json
{
  "url": "https://www.example.com/some/very/long/path"
}
```

**Example cURL:**

``` bash
curl -X POST http://localhost:8080/shorten   -H "Content-Type: application/json"   -d '{"url":"https://www.example.com/some/very/long/path"}'
```

**Example Response (new short URL):**

``` json
{
  "id": 1,
  "originalUrl": "https://www.example.com/some/very/long/path",
  "shortCode": "http://sho.rt/GS1OD0",
  "CreatedAt": "2025-02-07T10:00:00Z"
}
```

If the same URL is submitted again, the API **returns the existing short
URL instead of creating a new one.**


### **Redirect to original URL using short URL**

**Endpoint:**

    POST /:shortCode

**Example cURL:**

``` bash
curl -X GET http://sho.rt:8080/GS1OD0 
```

If the shortCode is valid, the API will redirect to the original URL

------------------------------------------------------------------------

## Database Behavior

-   Each `originalUrl` is **unique**
-   GORM auto-migrates the schema on startup
-   Duplicate URLs are prevented at both:
    -   Application level
    -   Database level (unique index)

Example model:

``` go
type ShortUrl struct {
    gorm.Model
    OriginalUrl string `gorm:"uniqueIndex"`
    ShortCode   string
}
```

------------------------------------------------------------------------

## Project Structure

    short-app/
    ├── main.go
    ├── db/
    │   └── database.go
    ├── model/
    │   └── shorten.go
    ├── handler/
    │   └── shorten_handler.go
    └── route/
        └── shorten_routes.go

------------------------------------------------------------------------

## How Short Codes Are Generated

-   Random alphanumeric string
-   Combined with `BASE_URL` from `.env`

------------------------------------------------------------------------

## Notes

-   No external API required
-   Designed as a minimal service
-   Can be deployed as a single Go binary

------------------------------------------------------------------------

## Future Improvements

-   Analytics (click count, timestamps)
-   Expiring links (TTL)
-   Custom aliases (e.g., `sho.rt/mylink`)
-   JWT authentication
-   Docker support
-   Frontend UI
