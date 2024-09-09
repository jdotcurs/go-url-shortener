package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/jdotcurs/go-url-shortener/internal/shortener"
)

type URLStore interface {
	Save(shortURL, longURL string) error
	Get(shortURL string) (string, bool)
}

// Handler handles HTTP requests for the URL shortener
type Handler struct {
	store   URLStore
	baseURL string
}

// NewHandler creates a new Handler
func NewHandler(store URLStore, baseURL string) *Handler {
	return &Handler{
		store:   store,
		baseURL: baseURL,
	}
}

// ShortenURL handles requests to shorten a URL
func (h *Handler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		LongURL string `json:"long_url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	shortURL := shortener.Shorten(req.LongURL)
	fullShortURL := shortener.GenerateShortURL(h.baseURL, req.LongURL)

	if err := h.store.Save(shortURL, req.LongURL); err != nil {
		http.Error(w, "Failed to save URL", http.StatusInternalServerError)
		return
	}

	resp := struct {
		ShortURL string `json:"short_url"`
	}{
		ShortURL: fullShortURL,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// RedirectURL handles redirection for short URLs
func (h *Handler) RedirectURL(w http.ResponseWriter, r *http.Request) {
	shortURL := strings.TrimPrefix(r.URL.Path, "/")
	longURL, exists := h.store.Get(shortURL)

	if !exists {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, longURL, http.StatusMovedPermanently)
}
