package requests

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/logan/v3/errors"

	"review_api/resources"
)

func NewUpdateRatingRequest(r *http.Request) (resources.UpdateRatingRequest, error) {
	var request resources.UpdateRatingRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal update rating request")
	}

	ratingIDStr := chi.URLParam(r, "rating_id")
	ratingId, err := strconv.ParseInt(ratingIDStr, 10, 64)
	if err != nil {
		return request, err
	}

	request.Data.RatingId = ratingId

	log.Printf("Decoded update rating request: %+v", request)
	return request, nil
}
