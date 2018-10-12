package main

import (
	"fmt"
	"net/http"
	"net/url"
)

// HTTPHandler - http requests handler
type HTTPHandler struct {
	urlMinifier URLMinifier
}

// NewHTTPHandler - return a new handler instance
func NewHTTPHandler(urlMinifier URLMinifier) *HTTPHandler {
	return &HTTPHandler{urlMinifier: urlMinifier}
}

// RedirectHandler - http hander: redirect request from minificated url to normal
func (h *HTTPHandler) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	path := url.PathEscape(r.URL.Path[1:])

	if redirectedPath := h.urlMinifier.GetFullURL(path); redirectedPath != "" {
		http.Redirect(w, r, redirectedPath, 301)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}

// MinificateHandler - http обработчик: минифицирует url, передаваемый в url параметре GET запроса
func (h *HTTPHandler) MinificateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if fullURL := r.URL.Query().Get("url"); fullURL != "" {
		minificatedURL := fmt.Sprintf("%s/%s", r.Host, h.urlMinifier.MinificateURL(fullURL))
		w.Write([]byte(minificatedURL))
		return
	}

	w.Write([]byte("No url param in request"))
	w.WriteHeader(http.StatusBadRequest)
}
