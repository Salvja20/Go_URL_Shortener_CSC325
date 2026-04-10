# Go URL Shortener — CSC325

A URL shortening service built in Go. Converts long URLs into short codes that redirect to the original URL. Built with a clean three-layer architecture: Storage, Service, and Handler.

---

## How to Run

**Requirements:**
- Go 1.26.2 or higher

**Start the server:**
```
go run main.go
```
Server will start at `http://localhost:8080`

---

## How to Use

### Shorten a URL
```powershell
Invoke-WebRequest -Uri "http://localhost:8080/shorten" -Method POST -ContentType "application/json" -Body '{"url": "https://github.com"}'
```

**Response:**
```json
{
  "original": "https://github.com",
  "short_code": "JEqMTO",
  "short_url": "http://localhost:8080/JEqMTO"
}
```

### Visit a Short URL
Open your browser and go to:
```
http://localhost:8080/JEqMTO
```
You will be redirected to the original URL.

---

## Project Structure

```
Go_URL_Shortener_CSC325/
  handlers/
    handler.go      ← HTTP endpoints, URL validation, request/response
  service/
    shortener.go    ← Business logic, short code generation
  storage/
    memory.go       ← Data storage, JSON persistence
  main.go           ← Entry point, server setup
  go.mod            ← Go module definition
  urls.json         ← Auto-generated, stores all shortened URLs
```

---

## Architecture

The project uses a three-layer architecture:

**Storage Layer** — Stores short code to URL mappings in memory and persists them to `urls.json` so data survives server restarts. Uses a `sync.RWMutex` for concurrency safety.

**Service Layer** — Contains all business logic. Generates unique 6-character short codes from letters and numbers. Ensures no two URLs share the same code.

**Handler Layer** — Handles HTTP requests. Validates URLs, calls the service, and returns JSON responses. Handles redirects and errors.

---

## Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/shorten` | Shorten a URL |
| GET | `/{code}` | Redirect to original URL |

---

## Error Handling

- Invalid URL format → `400 Bad Request`
- Short code not found → `404 Not Found`
- Wrong HTTP method → `405 Method Not Allowed`

---

## Built By
- Storage Layer — 
- Service Layer — 
- Handler Layer — 
