package main

import (
	"log"
	"net/http"

	"mini-web-service-go/internal/middlewares"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/items", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	handler := middlewares.RequestIDMiddleware(mux)
	handler = middlewares.LoggingMiddleware(handler)
	handler = middlewares.RecoveryMiddleware(handler)
	handler = middlewares.PerformanceMiddleware(handler)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
