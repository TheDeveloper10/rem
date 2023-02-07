package layer

import "net/http"

type userFriendlyError struct {
	Error string `json:"error"`
}

var (
	ErrStatusContentTypeNotJson = http.StatusUnsupportedMediaType

	ErrStatusNoBody = http.StatusBadRequest
	ErrMsgNoBody    = "No body"

	ErrStatusInvalidJson = http.StatusBadRequest
	ErrMsgInvalidJson    = "Invalid JSON"

	ErrStatusInvalidBody = http.StatusBadRequest

	ErrStatusInvalidURL = http.StatusBadRequest

	ErrStatusInvalidQuery = http.StatusBadRequest
)