package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"
)

// StorageService — Responsible for data storage
type StorageService struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewStorageService() *StorageService {
	return &StorageService{data: make(map[string]string)}
}

func (s *StorageService) Save(shortURL, longURL string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[shortURL] = longURL
}

func (s *StorageService) Get(shortURL string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	longURL, exists := s.data[shortURL]
	return longURL, exists
}

// URLShortener — Responsible for encoding and decoding
type URLShortener struct {
	storageService *StorageService
	baseURL         string
}

func NewURLShortener(storageService *StorageService, baseURL string) *URLShortener {
	return &URLShortener{storageService: storageService, baseURL: baseURL}
}

func (u *URLShortener) GenerateShortURL(longURL string) string {
	shortCode := generateRandomString(6)
	shortURL := fmt.Sprintf("%s/%s", u.baseURL, shortCode)
	u.storageService.Save(shortCode, longURL)
	return shortURL
}

func (u *URLShortener) ResolveShortURL(shortCode string) (string, bool) {
	return u.storageService.Get(shortCode)
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

// AnalyticsService — Responsible for tracking visits
type AnalyticsService struct {
	visitCount map[string]int
	mu         sync.Mutex
}

func NewAnalyticsService() *AnalyticsService {
	return &AnalyticsService{visitCount: make(map[string]int)}
}

func (a *AnalyticsService) RecordVisit(shortURL string) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.visitCount[shortURL]++
}

func (a *AnalyticsService) GetVisitCount(shortURL string) int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.visitCount[shortURL]
}

// HTTP Handlers for demonstrating functionality
func main() {
	storageService := NewStorageService()
	urlShortener := NewURLShortener(storageService, "http://short.url")
	analyticsService := NewAnalyticsService()

	// Example URLs
	longURL := "https://www.google.com/search?q=golang+system+design"
	shortURL := urlShortener.GenerateShortURL(longURL)

	fmt.Printf("Original URL: %s\n", longURL)
	fmt.Printf("Shortened URL: %s\n", shortURL)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		shortCode := strings.TrimPrefix(r.URL.Path, "/")
		longURL, found := urlShortener.ResolveShortURL(shortCode)
		if found {
			analyticsService.RecordVisit(shortCode)
			http.Redirect(w, r, longURL, http.StatusFound)
		} else {
			http.NotFound(w, r)
		}
	})

	http.HandleFunc("/analytics", func(w http.ResponseWriter, r *http.Request) {
		shortCode := r.URL.Query().Get("shortCode")
		count := analyticsService.GetVisitCount(shortCode)
		fmt.Fprintf(w, "URL %s has been visited %d times", shortCode, count)
	})

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}
