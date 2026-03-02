package middlewares

import (
	"log"
	"net/http"
	"time"

	"frameworks_first/internal/requestid"
)

func PerformanceMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)

		reqID := requestid.FromContext(r.Context())
		log.Printf("[PERF] Request %s: execution time %v ms", reqID, duration.Milliseconds())
	})
}
