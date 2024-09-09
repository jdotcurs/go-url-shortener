package shortener

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

// Shorten generates a short URL from a long URL
func Shorten(longURL string) string {
	hash := sha256.Sum256([]byte(longURL))
	return base64.URLEncoding.EncodeToString(hash[:8])
}

// GenerateShortURL creates the full short URL
func GenerateShortURL(baseURL, longURL string) string {
	return fmt.Sprintf("%s/%s", baseURL, Shorten(longURL))
}
