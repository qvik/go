package rest

import (
	"context"
	"net/http"
)

// Response is a generic REST API response
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

// Error defines the error type used for all errors from a RestHandler
type Error struct {
	Message    string
	StatusCode int
	Cause      error
}

// Error returns the error message
func (e *Error) Error() string {
	return e.Message
}

// Handler is our REST route handler type
type Handler func(ctx context.Context,
	r *http.Request) (responseData interface{}, err error)

// NewUnauthorizedError indicates the client/user was not authorized to
// perform this operation
func NewUnauthorizedError(message string) *Error {
	return &Error{StatusCode: http.StatusUnauthorized, Message: message}
}

// NewBadRequestError indicates the client sent an invalid request
func NewBadRequestError(message string) *Error {
	return &Error{StatusCode: http.StatusBadRequest, Message: message}
}

// NewNotFoundError indicates the resource that was the target of this
// request was not found
func NewNotFoundError(message string) *Error {
	return &Error{StatusCode: http.StatusNotFound, Message: message}
}

// NewForbiddenError indicates that it is forbidden to perform the request
func NewForbiddenError(message string) *Error {
	return &Error{StatusCode: http.StatusForbidden, Message: message}
}
