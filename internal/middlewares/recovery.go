package middlewares

import (
	"log"
	"net/http"
	"runtime/debug"

	"frameworks_first/internal/errors"
	"frameworks_first/internal/requestid"
)

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				reqID := requestid.FromContext(r.Context())
				log.Printf("[PANIC] Request %s: %v\n%s", reqID, rec, debug.Stack())
				errors.HandleError(w, r, errors.ErrInternal)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
