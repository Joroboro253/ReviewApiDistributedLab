package resources

import (
	"net/http"
)

type APIError struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

func NewAPIError(status int, title, detail string) *APIError {
	return &APIError{
		Status: status,
		Title:  title,
		Detail: detail,
	}

}

func (e *APIError) Error() string {
	return e.Title + ": " + e.Detail
}

var (
	ErrReviewNotFound  = NewAPIError(http.StatusNotFound, "Review Not Found", "The requested review does not exist")
	ErrInvalidInput    = NewAPIError(http.StatusBadRequest, "Invalid Input", "The provided input is not valid")
	ErrDatabaseProblem = NewAPIError(http.StatusInternalServerError, "Database Problem", "A problem occurred with the database")
	ErrInternal        = NewAPIError(http.StatusInternalServerError, "Internal Error", "An internal error occurred")
)
