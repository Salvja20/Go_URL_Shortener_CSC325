package gourlshortenercsc325

import (
	"fmt"
	"net/http"
	"Go_URL_Shortener_CSC325/handlers"
)

func main() {
	http.HandleFunc("/shorten", handlers.ShortenURL)
	http.HandleFunc("/", handlers.RedirectURL)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
