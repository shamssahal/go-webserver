package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Error struct {
	Code int    `json:"code"`
	Err  string `json:"error"`
}

func NewError(code int, err string) Error {
	return Error{
		Code: code,
		Err:  err,
	}
}

func ErrUnauthorized() Error {
	return Error{
		Code: http.StatusUnauthorized,
		Err:  "unauthorized request",
	}
}

func ErrTokenExpired() Error {
	return Error{
		Code: http.StatusUnauthorized,
		Err:  "token expired",
	}
}

func ErrBadRequest() Error {
	return Error{
		Code: http.StatusBadRequest,
		Err:  "invalid JSON request",
	}
}

func ErrResourceNotFound(res string) Error {
	return Error{
		Code: http.StatusNotFound,
		Err:  fmt.Sprintf("%s resource not found", res),
	}
}

// Error implements the error interface
func (e Error) Error() string {
	return e.Err
}

// StatusCode returns the HTTP status code
func (e Error) StatusCode() int {
	return e.Code
}

// WriteError writes an error response in JSON format
func WriteError(w http.ResponseWriter, err Error) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Code)

	response := map[string]any{
		"error": err.Err,
		"code":  err.Code,
	}

	return json.NewEncoder(w).Encode(response)
}
