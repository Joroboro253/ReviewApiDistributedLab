package requests

import (
	"encoding/json"
	"net/http"
	"review_api/resources"
)

// Здесь происходит сам запрос. Нужно проверить на роботоспособность
func DecodeReviewRequestBody(r *http.Request) (resources.ReviewRequest, *resources.APIError) {
	var reqBody resources.ReviewRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		return reqBody, resources.NewAPIError(http.StatusBadRequest, "Error during JSON decoding", err.Error())
	}
	return reqBody, nil
}
