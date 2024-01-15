package requests

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type RateAttributes struct {
	ReviewID *int64   `json:"review_id"`
	UserID   *int64   `json:"user_id"`
	Rating   *float64 `json:"rating"`
}

type UpdateRatingData struct {
	Attributes RateAttributes `json:"attributes"`
}

type UpdateRatingRequest struct {
	Data     UpdateRatingData `json:"data"`
	RatingID int64
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

	log.Printf("Decoded update rating request: %+v", request)
	return request, nil
}
