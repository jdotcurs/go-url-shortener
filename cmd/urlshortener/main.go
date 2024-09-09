package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/jdotcurs/go-url-shortener/internal/handler"
	"github.com/jdotcurs/go-url-shortener/internal/store"
)

func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func main() {
	port := flag.Int("port", 8080, "Port to run the server on")
	baseURL := flag.String("base-url", "http://localhost:8080", "Base URL for shortened URLs")
	flag.Parse()

	urlStore := store.NewURLStore()
	h := handler.NewHandler(urlStore, *baseURL)

	http.HandleFunc("/shorten", enableCORS(h.ShortenURL))
	http.HandleFunc("/", enableCORS(h.RedirectURL))

	addr := fmt.Sprintf(":%d", *port)
	log.Printf("Starting server on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
