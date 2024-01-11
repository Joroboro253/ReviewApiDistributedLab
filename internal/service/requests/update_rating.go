package requests

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type UpdateRatingRequest struct {
	Data struct {
		ReviewID *int64   `json:"rating_id"`
		UserID   *int64   `json:"user_id"`
		Rating   *float64 `json:"rating"`
	} `json:"data"`
	RatingID int64 `json:"-"`
}

func NewUpdateRatingRequest(r *http.Request) (UpdateRatingRequest, error) {
	var request UpdateRatingRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal update rating request")
	}

	ratingIDStr := chi.URLParam(r, "rating_id")
	ratingID, err := strconv.ParseInt(ratingIDStr, 10, 64)
	if err != nil {
		return request, err
	}

	request.RatingID = ratingID

	return request, nil
}
