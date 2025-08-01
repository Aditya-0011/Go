package middlewares

import (
	"log"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rr *responseWriter) WriteHeader(code int) {
	rr.statusCode = code
	rr.ResponseWriter.WriteHeader(code)
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rr := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(rr, r)

		duration := time.Since(start)

		log.Printf("[%s] %s %d %s",
			r.Method,
			r.URL.Path,
			rr.statusCode,
			duration,
		)
	})
}
