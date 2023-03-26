# Movies API
Movies API (CRUD) using Go, GORM, PostgreSQL, and Gorilla Mux for routing.

## Requirements
1. GO >= 1.20
2. PostgreSQL >= 12

## How to run this app locally ?
Follow these steps to run this application locally.

### Step 1
Move into /pkg/config directory and open the "app.go" file.

```
dsn := "host=localhost user=postgres password=78678612 dbname=gomoviesapi port=5432 sslmode=disable TimeZone=Asia/Karachi"
```
*Change the following details with your's details.*
```
dsn := "host=localhost user=your user name password=your password dbname=your db name port=5432 sslmode=disable TimeZone=your timezone"
```

### Step 2
After changing the details move into cmd/main directory and run the following command.

```
go run main.go
```
