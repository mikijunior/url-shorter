package storage

import (
	"errors"
	"sync"
	"url_shortener/utils"
)

var (
	urls = make(map[string]string)
	mu   sync.RWMutex
)

func SaveURL(originalURL string) string {
	mu.Lock()
	defer mu.Unlock()

	shortURL := utils.GenerateShortURL()
	urls[shortURL] = originalURL
	return shortURL
}

func GetURL(shortURL string) (string, error) {
	mu.RLock()
	defer mu.RUnlock()

	originalURL, exists := urls[shortURL]
	if !exists {
		return "", errors.New("URL not found")
	}

	return originalURL, nil
}
