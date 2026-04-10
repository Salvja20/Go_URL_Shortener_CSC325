package handlers
 
import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
 
	"Go_URL_Shortener_CSC325/service"
)

var shortenerService = service.NewShortenerService()
 
func ShortenURL(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed. Use POST.")
		return
	}
 
	
	var request struct {
		URL string `json:"url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid JSON body.")
		return
	}
 
	
	if err := validateURL(request.URL); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
 

	code := shortenerService.Shorten(request.URL)
 
	
	shortURL := "http://localhost:8080/" + code
 
	respondWithJSON(w, http.StatusOK, map[string]string{
		"short_code": code,
		"short_url":  shortURL,
		"original":   request.URL,
	})
}
 

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed. Use GET.")
		return
	}
 
	code := strings.TrimPrefix(r.URL.Path, "/")
 
	
	if code == "" {
		respondWithJSON(w, http.StatusOK, map[string]string{
			"message": "URL Shortener is running. POST to /shorten to create a short URL.",
		})
		return
	}
 
	
	originalURL, exists := shortenerService.GetOriginalURL(code)
	if !exists {
		respondWithError(w, http.StatusNotFound, "Short URL not found: "+code)
		return
	}
 
	
	http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
}
 
func validateURL(rawURL string) error {
	
	if strings.TrimSpace(rawURL) == "" {
		return &ValidationError{"URL cannot be empty."}
	}
 
	
	parsed, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return &ValidationError{"Invalid URL format. Must include http:// or https://"}
	}
 
	
	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return &ValidationError{"URL must start with http:// or https://"}
	}
 
	
	if parsed.Host == "" {
		return &ValidationError{"URL must include a valid domain."}
	}
 
	return nil
}
 

type ValidationError struct {
	Message string
}
 
func (e *ValidationError) Error() string {
	return e.Message
}
 

func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
 

func respondWithError(w http.ResponseWriter, status int, message string) {
	respondWithJSON(w, status, map[string]string{
		"error": message,
	})
}
