package helpers

import (
	"encoding/json"
	"gitlab.com/distributed_lab/kit/pgdb"
	"net/http"
	"review_api/resources"
)

type GetHandler struct {
	DB *pgdb.DB
}

func (h *GetHandler) QueryParameterProcessing(r *http.Request) (string, string, string) {
	sortField := r.URL.Query().Get("sort")
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")
	return sortField, pageStr, limitStr
}

func (h *GetHandler) GenerateResponse(w http.ResponseWriter, response map[string]interface{}) *resources.APIError {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		return resources.NewAPIError(http.StatusInternalServerError, "StatusInternalServerError", "Error encoding response")
	}
	return nil
}
