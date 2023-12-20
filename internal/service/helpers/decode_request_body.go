package helpers

import (
	"encoding/json"
	"log"
	"net/http"
	"review_api/resources"
)

func (h *CreateHandler) DecodeRequestBody(r *http.Request) (resources.UpdateRequest, *resources.APIError) {
	var reqBody resources.UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		errorMsg := "Error during JSON decoding"
		log.Printf("%s: %v", errorMsg, err)
		return resources.UpdateRequest{}, resources.NewAPIError(http.StatusBadRequest, "StatusBadRequest", errorMsg)
	}
	return reqBody, nil
}
