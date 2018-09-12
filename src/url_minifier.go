package main

import (
	"hash/adler32"
	"strconv"
)

// MinificateURL - минифицирует url и возвращает url path для его которкой версии
func MinificateURL(store map[string]string, longURL string) string {
	hash := computeURLHash(longURL)

	// Adler-32 имеет коллизии, поэтому нужно перезаписывать значение в хранилище
	store[hash] = longURL
	return hash
}

// GetLongURL - возвращает полный url по минифицированному
func GetLongURL(store map[string]string, shortURL string) string {
	if longURL, ok := store[shortURL]; ok {
		return longURL
	}

	return ""
}

func computeURLHash(url string) string {
	adlerHash := adler32.Checksum([]byte(url))
	return strconv.FormatUint(uint64(adlerHash), 16)
}
