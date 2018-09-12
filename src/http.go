package main

import (
	"fmt"
	"net/http"
	"net/url"
)

// HTTPHandler - обработчик http запросов
type HTTPHandler struct {
	urlStore map[string]string
}

// NewHTTPHandler - создает новый инстанс обработчика запросов
func NewHTTPHandler() *HTTPHandler {
	urlStore := make(map[string]string)
	return &HTTPHandler{urlStore}
}

// RedirectHandler - http обработчик: перенаправляет запрос с минифицированного url на нормальный
func (h HTTPHandler) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	path := url.PathEscape(r.URL.Path[1:])

	if redirectedPath := GetLongURL(h.urlStore, path); redirectedPath != "" {
		http.Redirect(w, r, redirectedPath, 301)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}

// MinificateHandler - http обработчик: минифицирует url, передаваемый в url параметре GET запроса
func (h HTTPHandler) MinificateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if longURL := r.URL.Query().Get("url"); longURL != "" {
		minificatedURL := fmt.Sprintf("%s/%s", r.Host, MinificateURL(h.urlStore, longURL))
		w.Write([]byte(minificatedURL))
		return
	}

	w.Write([]byte("No url param in request"))
}
