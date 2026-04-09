package service

import (
	"math/rand"
	"time"
	"Go_URL_Shortener_CSC325/storage"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type ShortenerService struct {
	store *storage.MemoryStore
}

func NewShortenerService() *ShortenerService {
	return &ShortenerService{
		store: storage.NewMemoryStore(),
	}
}

// Generate random short code
func generateCode(length int) string {
	rand.Seed(time.Now().UnixNano())
	code := make([]byte, length)

	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}

	return string(code)
}

func (s *ShortenerService) Shorten(url string) string {
	code := generateCode(6)

	// Ensure uniqueness
	for {
		_, exists := s.store.Get(code)
		if !exists {
			break
		}
		code = generateCode(6)
	}

	s.store.Save(code, url)
	return code
}

func (s *ShortenerService) GetOriginalURL(code string) (string, bool) {
	return s.store.Get(code)
}