package store

import (
	"sync"
)

// URLStore is an in-memory store for URL mappings
type URLStore struct {
	shortToLong map[string]string
	longToShort map[string]string
	mutex       sync.RWMutex
}

// NewURLStore creates a new URLStore
func NewURLStore() *URLStore {
	return &URLStore{
		shortToLong: make(map[string]string),
		longToShort: make(map[string]string),
	}
}

// Save stores a short URL and its corresponding long URL
func (s *URLStore) Save(shortURL, longURL string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.shortToLong[shortURL] = longURL
	s.longToShort[longURL] = shortURL
	return nil
}

// Get retrieves the long URL for a given short URL
func (s *URLStore) Get(shortURL string) (string, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	longURL, exists := s.shortToLong[shortURL]
	return longURL, exists
}

// GetShortURL retrieves the short URL for a given long URL
func (s *URLStore) GetShortURL(longURL string) (string, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	shortURL, exists := s.longToShort[longURL]
	return shortURL, exists
}
