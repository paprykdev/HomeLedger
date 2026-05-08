package middleware

import (
	"log"
	"net/http"
	"time"
)

type StatusRecorder struct {
	http.ResponseWriter
	StatusCode int
}

func (r *StatusRecorder) WriteHeader(statusCode int) {
	r.StatusCode = statusCode

	r.ResponseWriter.WriteHeader(statusCode)
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		start := time.Now()

		recorder := &StatusRecorder{
			ResponseWriter: w,
			StatusCode: 200,
		}
		defer func() {
			log.Printf(
				"[%s] %s: %s %d %dms",
				time.Now().Format(time.DateTime),
				r.Method,
				r.URL.Path,
				recorder.StatusCode,
				time.Since(start).Milliseconds(),
			)
		}()

		next.ServeHTTP(recorder, r)

	})
}
