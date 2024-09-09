package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	// Save original os.Args
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	// Set up test arguments
	os.Args = []string{"cmd", "-port", "8080", "-base-url", "http://test.com"}

	// Create a channel to signal when the server is ready
	ready := make(chan bool)

	// Start the server in a goroutine
	go func() {

		log.SetOutput(logWriter{ready: ready})
		defer log.SetOutput(os.Stderr)

		main()
	}()

	// Wait for the server to start or timeout
	select {
	case <-ready:
		// Server started successfully
	case <-time.After(5 * time.Second):
		t.Fatal("Server didn't start within the expected time")
	}

	// Make a test request
	resp, err := http.Get("http://localhost:8080/")
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", resp.StatusCode)
	}
}

type logWriter struct {
	ready chan<- bool
}

func (w logWriter) Write(p []byte) (int, error) {
	w.ready <- true
	return len(p), nil
}

func TestCORSPreflight(t *testing.T) {
	req, err := http.NewRequest("OPTIONS", "http://localhost:8080/shorten", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enableCORS(func(w http.ResponseWriter, r *http.Request) {})(w, r)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code for OPTIONS: got %v want %v", status, http.StatusOK)
	}

	expectedHeaders := map[string]string{
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Methods": "POST, GET, OPTIONS",
		"Access-Control-Allow-Headers": "Content-Type",
	}

	for key, value := range expectedHeaders {
		if rr.Header().Get(key) != value {
			t.Errorf("Handler returned wrong header for %s: got %v want %v", key, rr.Header().Get(key), value)
		}
	}
}
