package config

import (
	"server/models"
	"sync"
	"time"
)

var (
	URLStore = &URLStorage{
		urls: make(map[string]*models.Url),
	}
)

type URLStorage struct {
	urls map[string]*models.Url
	mu   sync.RWMutex
}

// GetURL retrieves a URL by short code
func (s *URLStorage) GetURL(shortCode string) *models.Url {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.urls[shortCode]
}

// CreateURL creates a new shortened URL
func (s *URLStorage) CreateURL(shortCode string, originalUrl string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.urls[shortCode]; exists {
		return ErrShortCodeExists
	}

	s.urls[shortCode] = &models.Url{
		ShortCode:   shortCode,
		OriginalUrl: originalUrl,
		CreatedAt:   time.Now(),
		ExpiresAt:   time.Now().Add(30 * 24 * time.Hour), // 30 days
		Clicks:      0,
	}

	return nil
}

// DeleteURL deletes a URL by short code
func (s *URLStorage) DeleteURL(shortCode string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.urls[shortCode]; !exists {
		return ErrURLNotFound
	}

	delete(s.urls, shortCode)
	return nil
}

// CheckAvailability checks if a short code is available
func (s *URLStorage) CheckAvailability(shortCode string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	_, exists := s.urls[shortCode]
	return !exists
}

// IncrementClicks increments click count for a URL
func (s *URLStorage) IncrementClicks(shortCode string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	url, exists := s.urls[shortCode]
	if !exists {
		return ErrURLNotFound
	}

	url.Clicks++
	return nil
}

// GetAllURLs gets all stored URLs (for debugging/stats)
func (s *URLStorage) GetAllURLs() []*models.Url {
	s.mu.RLock()
	defer s.mu.RUnlock()

	urls := make([]*models.Url, 0, len(s.urls))
	for _, url := range s.urls {
		urls = append(urls, url)
	}
	return urls
}
