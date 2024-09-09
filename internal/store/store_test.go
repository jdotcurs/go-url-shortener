package store

import (
	"testing"
)

func TestURLStore(t *testing.T) {
	store := NewURLStore()

	store.Save("abc123", "https://example.com")
	longURL, exists := store.Get("abc123")
	if !exists {
		t.Errorf("Expected long URL to exist, but it does not")
	}
	if longURL != "https://example.com" {
		t.Errorf("Expected long URL to be 'https://example.com', but got '%s'", longURL)
	}

	store.Save("abc123", "https://another-example.com")
	longURL, exists = store.Get("abc123")
	if !exists {
		t.Errorf("Expected long URL to exist, but it does not")
	}
	if longURL != "https://another-example.com" {
		t.Errorf("Expected long URL to be 'https://another-example.com', but got '%s'", longURL)
	}
}

func TestURLStore_GetShortURL(t *testing.T) {
	store := NewURLStore()

	longURL := "https://example.com"
	shortURL := "abc123"
	store.Save(shortURL, longURL)

	gotShortURL, exists := store.GetShortURL(longURL)
	if !exists {
		t.Errorf("Expected short URL to exist, but it does not")
	}
	if gotShortURL != shortURL {
		t.Errorf("Expected short URL to be '%s', but got '%s'", shortURL, gotShortURL)
	}

	_, exists = store.GetShortURL("https://nonexistent.com")
	if exists {
		t.Errorf("Expected non-existent URL to return false, but it returned true")
	}
}
