package config

import "errors"

var (
	ErrShortCodeExists = errors.New("short code already exists")
	ErrURLNotFound     = errors.New("URL not found")
	ErrInvalidURL      = errors.New("invalid URL format")
)
