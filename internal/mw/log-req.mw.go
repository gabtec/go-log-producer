package mw

import (
	"log"
	"net/http"
	"time"
)

// ReqLogger logs each incoming HTTP request with method, path, and duration.
func ReqLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("%s %s __ Completed in %v", r.Method, r.URL.Path, time.Since(start))
	})
}
