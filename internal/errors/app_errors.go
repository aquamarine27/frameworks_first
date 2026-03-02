package errors

import (
	"encoding/json"
	"net/http"

	"mini-web-service-go/internal/requestid"
)

type AppError struct {
	Code    string
	Message string
	Status  int
}

func (e AppError) Error() string {
	return e.Message
}

func (e AppError) WithMessage(msg string) AppError {
	e.Message = msg
	return e
}

type ErrorResponse struct {
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
	RequestID string `json:"requestId"`
}

var (
	ErrNotFound    = AppError{"NOT_FOUND", "Item not found", http.StatusNotFound}
	ErrInvalidID   = AppError{"INVALID_ID", "Invalid ID", http.StatusBadRequest}
	ErrInvalidJSON = AppError{"INVALID_JSON", "Invalid request body", http.StatusBadRequest}
	ErrValidation  = AppError{"VALIDATION_ERROR", "Validation failed", http.StatusBadRequest}
	ErrInternal    = AppError{"INTERNAL_ERROR", "Internal server error", http.StatusInternalServerError}
)

func HandleError(w http.ResponseWriter, r *http.Request, err error) {
	appErr := ErrInternal
	if ae, ok := err.(AppError); ok {
		appErr = ae
	}

	reqID := requestid.FromContext(r.Context())

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(appErr.Status)
	_ = json.NewEncoder(w).Encode(ErrorResponse{
		ErrorCode: appErr.Code,
		Message:   appErr.Message,
		RequestID: reqID,
	})
}

func RespondJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}
