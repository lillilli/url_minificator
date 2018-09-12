package main

import (
	"log"
	"net/http"
	"time"
)

const (
	serverAddr   = ":8080"
	readTimeout  = time.Duration(1 * time.Second)
	writeTimeout = readTimeout
)

func main() {
	handler := NewHTTPHandler()

	http.HandleFunc("/", handler.RedirectHandler)
	http.HandleFunc("/--", handler.MinificateHandler)

	srv := http.Server{
		Addr:         serverAddr,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	log.Printf("Starting server on %s", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
