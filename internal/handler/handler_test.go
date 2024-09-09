package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jdotcurs/go-url-shortener/internal/store"
)

type mockURLStore struct {
	saveFunc func(shortURL, longURL string) error
	getFunc  func(shortURL string) (string, bool)
}

func (m *mockURLStore) Save(shortURL, longURL string) error {
	return m.saveFunc(shortURL, longURL)
}

func (m *mockURLStore) Get(shortURL string) (string, bool) {
	return m.getFunc(shortURL)
}

func TestHandler_ShortenURL(t *testing.T) {
	urlStore := &mockURLStore{
		saveFunc: func(shortURL, longURL string) error {
			if longURL == "https://error.com" {
				return fmt.Errorf("mock error")
			}
			return nil
		},
	}
	handler := NewHandler(urlStore, "http://short.url")

	tests := []struct {
		name           string
		method         string
		body           map[string]string
		expectedStatus int
	}{
		{"Valid request", http.MethodPost, map[string]string{"long_url": "https://example.com"}, http.StatusOK},
		{"Invalid method", http.MethodGet, nil, http.StatusMethodNotAllowed},
		{"Invalid body", http.MethodPost, nil, http.StatusBadRequest},
		{"Store save error", http.MethodPost, map[string]string{"long_url": "https://error.com"}, http.StatusInternalServerError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var body []byte
			var err error
			if tt.body != nil {
				body, err = json.Marshal(tt.body)
				if err != nil {
					t.Fatalf("Failed to marshal request body: %v", err)
				}
			}

			req, err := http.NewRequest(tt.method, "/shorten", bytes.NewBuffer(body))
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			rr := httptest.NewRecorder()
			handler.ShortenURL(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("Handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}

			if tt.expectedStatus == http.StatusOK {
				var response map[string]string
				err = json.Unmarshal(rr.Body.Bytes(), &response)
				if err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if _, ok := response["short_url"]; !ok {
					t.Errorf("Response does not contain short_url")
				}
			}
		})
	}
}

func TestHandler_RedirectURL(t *testing.T) {
	urlStore := store.NewURLStore()
	handler := NewHandler(urlStore, "http://short.url")

	// Add a test URL to the store
	shortURL := "testshort"
	longURL := "https://example.com"
	urlStore.Save(shortURL, longURL)

	tests := []struct {
		name           string
		shortURL       string
		expectedStatus int
		expectedURL    string
	}{
		{"Valid short URL", shortURL, http.StatusMovedPermanently, longURL},
		{"Invalid short URL", "nonexistent", http.StatusNotFound, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, "/"+tt.shortURL, nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			rr := httptest.NewRecorder()
			handler.RedirectURL(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("Handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}

			if tt.expectedStatus == http.StatusMovedPermanently {
				location := rr.Header().Get("Location")
				if location != tt.expectedURL {
					t.Errorf("Handler returned wrong redirect URL: got %v want %v", location, tt.expectedURL)
				}
			}
		})
	}
}
