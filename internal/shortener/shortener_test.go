package shortener

import (
	"testing"
)

func TestShorten(t *testing.T) {
	tests := []struct {
		name     string
		longURL  string
		expected int
	}{
		{"Simple URL", "https://example.com", 12},
		{"Empty string", "", 12},
		{"Long URL", "https://example.com/very/long/url/with/many/parameters?param1=value1&param2=value2", 12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Shorten(tt.longURL)
			if len(result) != tt.expected {
				t.Errorf("Shorten() = %v, want length %v", result, tt.expected)
			}
		})
	}
}

func TestGenerateShortURL(t *testing.T) {
	tests := []struct {
		name     string
		baseURL  string
		longURL  string
		expected string
	}{
		{"Simple URL", "http://short.url", "https://example.com", "http://short.url/"},
		{"Base URL with path", "http://short.url/api", "https://example.com", "http://short.url/api/"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GenerateShortURL(tt.baseURL, tt.longURL)
			if len(result) <= len(tt.expected) {
				t.Errorf("GenerateShortURL() = %v, expected longer than %v", result, tt.expected)
			}
			if result[:len(tt.expected)] != tt.expected {
				t.Errorf("GenerateShortURL() = %v, expected to start with %v", result, tt.expected)
			}
		})
	}
}
