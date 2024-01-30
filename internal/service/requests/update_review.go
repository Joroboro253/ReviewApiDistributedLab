package requests

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/logan/v3/errors"

	"review_api/resources"
)

func NewUpdateReviewRequest(r *http.Request) (resources.UpdateReviewRequest, error) {
	var request resources.UpdateReviewRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal update review request")
	}

	reviewIDStr := chi.URLParam(r, "review_id")
	reviewId, err := strconv.ParseInt(reviewIDStr, 10, 64)
	if err != nil {
		return request, err
	}

	request.Data.ReviewId = reviewId

	return request, nil
}
