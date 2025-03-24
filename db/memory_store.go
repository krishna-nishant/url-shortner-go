package db

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"
)

// URL represents a shortened URL in our system
type URL struct {
	ID         uint      `json:"id"`
	Original   string    `json:"original"`
	Short      string    `json:"short"`
	ClickCount int       `json:"click_count"`
	CreatedAt  time.Time `json:"created_at"`
}

// MemoryStore provides an in-memory implementation of the URL store
// with optional persistence to a JSON file
type MemoryStore struct {
	urls      map[string]URL
	urlMutex  *sync.RWMutex
	lastID    uint
	storePath string
}

// NewMemoryStore creates a new MemoryStore
func NewMemoryStore(persistToFile bool) *MemoryStore {
	store := &MemoryStore{
		urls:     make(map[string]URL),
		urlMutex: &sync.RWMutex{},
		lastID:   0,
	}

	if persistToFile {
		store.storePath = "./db/urls.json"
		store.loadFromFile()

		// Create a goroutine that periodically saves data to disk
		go func() {
			ticker := time.NewTicker(1 * time.Minute)
			for range ticker.C {
				store.saveToFile()
			}
		}()
	}

	return store
}

// SaveURL saves a URL to the store
func (s *MemoryStore) SaveURL(url URL) URL {
	s.urlMutex.Lock()
	defer s.urlMutex.Unlock()

	// If this is a new URL, assign an ID
	if url.ID == 0 {
		s.lastID++
		url.ID = s.lastID
		url.CreatedAt = time.Now()
	}

	s.urls[url.Short] = url
	return url
}

// GetByShortURL retrieves a URL by its short code
func (s *MemoryStore) GetByShortURL(shortURL string) (URL, bool) {
	s.urlMutex.RLock()
	defer s.urlMutex.RUnlock()

	url, exists := s.urls[shortURL]
	return url, exists
}

// UpdateClickCount increments the click count for a URL
func (s *MemoryStore) UpdateClickCount(shortURL string) {
	s.urlMutex.Lock()
	defer s.urlMutex.Unlock()

	if url, exists := s.urls[shortURL]; exists {
		url.ClickCount++
		s.urls[shortURL] = url
	}
}

// GetAllURLs returns all URLs in the store
func (s *MemoryStore) GetAllURLs() []URL {
	s.urlMutex.RLock()
	defer s.urlMutex.RUnlock()

	urlList := make([]URL, 0, len(s.urls))
	for _, url := range s.urls {
		urlList = append(urlList, url)
	}

	return urlList
}

// loadFromFile loads URLs from a JSON file if it exists
func (s *MemoryStore) loadFromFile() {
	if _, err := os.Stat(s.storePath); os.IsNotExist(err) {
		// File doesn't exist yet, that's OK
		return
	}

	data, err := ioutil.ReadFile(s.storePath)
	if err != nil {
		log.Printf("Error reading URL data from file: %v", err)
		return
	}

	var storedData struct {
		LastID uint           `json:"last_id"`
		URLs   map[string]URL `json:"urls"`
	}

	if err := json.Unmarshal(data, &storedData); err != nil {
		log.Printf("Error parsing URL data from file: %v", err)
		return
	}

	s.urlMutex.Lock()
	defer s.urlMutex.Unlock()

	s.urls = storedData.URLs
	s.lastID = storedData.LastID

	log.Printf("Loaded %d URLs from storage", len(s.urls))
}

// saveToFile saves URLs to a JSON file
func (s *MemoryStore) saveToFile() {
	// Ensure directory exists
	dir := "./db"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.Mkdir(dir, 0755); err != nil {
			log.Printf("Error creating directory: %v", err)
			return
		}
	}

	s.urlMutex.RLock()
	storedData := struct {
		LastID uint           `json:"last_id"`
		URLs   map[string]URL `json:"urls"`
	}{
		LastID: s.lastID,
		URLs:   s.urls,
	}
	s.urlMutex.RUnlock()

	data, err := json.MarshalIndent(storedData, "", "  ")
	if err != nil {
		log.Printf("Error serializing URL data: %v", err)
		return
	}

	if err := ioutil.WriteFile(s.storePath, data, 0644); err != nil {
		log.Printf("Error writing URL data to file: %v", err)
		return
	}

	log.Printf("Saved %d URLs to storage", len(s.urls))
}
