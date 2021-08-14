package apierrors

import (
	"fmt"
	"net/http"
)

type ApiErr interface {
	Message() string
	Status() int
	Error() string
}

type apiErr struct {
	ErrMessage string `json:"message"`
	ErrStatus  int    `json:"status"`
	Err        string `json:"error"`
}

func (e apiErr) Error() string {
	return fmt.Sprintf("message: %s - status: %d error: %s", e.ErrMessage, e.ErrStatus, e.Err)
}

func (e apiErr) Message() string {
	return e.ErrMessage
}

func (e apiErr) Status() int {
	return e.ErrStatus
}

func NewBadRequestError(message string) ApiErr {
	return apiErr{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		Err:        "bad_request",
	}
}

func NewNotFoundError(message string) ApiErr {
	return apiErr{
		ErrMessage: message,
		ErrStatus:  http.StatusNotFound,
		Err:        "not_found",
	}
}

func NewInternalServerError(message string) ApiErr {
	return apiErr{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		Err:        "internal_server_error",
	}
}
