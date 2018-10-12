package main

import (
	"hash/adler32"
	"strconv"
	"sync"
)

// URLMinifier -url minifier interface
type URLMinifier interface {
	MinificateURL(longURL string) string
	GetFullURL(urlHash string) string
}

type urlMinifier struct {
	store map[string]string
	sync.Mutex
}

// NewURLMinifer - return a new url minifier
func NewURLMinifer() URLMinifier {
	return &urlMinifier{store: make(map[string]string)}
}

// MinificateURL - minificate url and return url hash
func (m *urlMinifier) MinificateURL(fullURL string) string {
	hash := computeURLHash(fullURL)

	// Adler-32 have a collisions, that`s why we need to overwrite value in store
	m.Lock()
	m.store[hash] = fullURL
	m.Unlock()

	return hash
}

// GetLongURL - return a full url from a url hash
func (m *urlMinifier) GetFullURL(urlHash string) string {
	if fullURL, ok := m.store[urlHash]; ok {
		return fullURL
	}

	return ""
}

func computeURLHash(url string) string {
	adlerHash := adler32.Checksum([]byte(url))
	return strconv.FormatUint(uint64(adlerHash), 16)
}
