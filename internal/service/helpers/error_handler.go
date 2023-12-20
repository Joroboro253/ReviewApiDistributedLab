package helpers

import (
	"net/http"
	"review_api/resources"
)

func ErrorHandler(handler func(http.ResponseWriter, *http.Request) *resources.APIError) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if apiErr := handler(w, r); apiErr != nil {
			SendApiError(w, apiErr)
		}
	}
}
