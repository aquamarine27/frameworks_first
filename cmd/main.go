package main

import (
	"log"
	"net/http"

	"frameworks_first/internal/domain"
	"frameworks_first/internal/errors"
	"frameworks_first/internal/middlewares"
	"frameworks_first/internal/services"

	"encoding/json"
	"strconv"
)

func main() {
	repo := services.NewInMemoryRepository()
	service := services.NewTaskService(repo)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/items", func(w http.ResponseWriter, r *http.Request) {
		items, err := service.GetAll()
		if err != nil {
			errors.HandleError(w, r, err)
			return
		}
		errors.RespondJSON(w, http.StatusOK, items)
	})

	mux.HandleFunc("GET /api/items/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			errors.HandleError(w, r, errors.ErrInvalidID)
			return
		}
		item, err := service.GetByID(id)
		if err != nil {
			errors.HandleError(w, r, err)
			return
		}
		errors.RespondJSON(w, http.StatusOK, item)
	})

	mux.HandleFunc("POST /api/items", func(w http.ResponseWriter, r *http.Request) {
		var item domain.TaskItem
		if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
			errors.HandleError(w, r, errors.ErrInvalidJSON)
			return
		}
		newItem, err := service.Create(&item)
		if err != nil {
			errors.HandleError(w, r, err)
			return
		}
		errors.RespondJSON(w, http.StatusCreated, newItem)
	})

	// Конвейер middleware
	handler := middlewares.RequestIDMiddleware(mux)
	handler = middlewares.LoggingMiddleware(handler)
	handler = middlewares.RecoveryMiddleware(handler)
	handler = middlewares.PerformanceMiddleware(handler)

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
