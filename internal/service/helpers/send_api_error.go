package helpers

import (
	"encoding/json"
	"net/http"
	"review_api/resources"
)

func SendApiError(w http.ResponseWriter, apiErr *resources.APIError) {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(apiErr.Status)
	json.NewEncoder(w).Encode(map[string][]resources.APIError{"errors": {*apiErr}})
}
