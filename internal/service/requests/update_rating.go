package requests

import (
	"encoding/json"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
)

type UpdateRatingRequest struct {
	Data struct {
		ReviewID int64   `json:"rating_id"`
		UserID   int64   `json:"user_id"`
		Rating   float64 `json:"rating"`
	} `json:"data"`
}

func NewUpdateRatingRequest(r *http.Request) (UpdateRatingRequest, error) {
	var request UpdateRatingRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal update rating request")
	}
	return request, nil
}
