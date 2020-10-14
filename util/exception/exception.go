package exception

import "net/http"

type Exception struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"exception"`
}

func NewException(message string, status int, error string) *Exception {
	return &Exception{
		Message: message,
		Status:  status,
		Error:   error,
	}
}

func NewInternalErrorException(message string) *Exception {
	return NewException(message, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
}

func NewBadRequestException(message string) *Exception {
	return NewException(message, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
}

func NewNotFoundException(message string) *Exception {
	return NewException(message, http.StatusNotFound, http.StatusText(http.StatusNotFound))
}
