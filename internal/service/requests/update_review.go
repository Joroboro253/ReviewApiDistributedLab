package requests

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type RevAttributes struct {
	ProductID *int64  `json:"product_id"`
	UserID    *int64  `json:"user_id"`
	Content   *string `json:"content"`
}

type UpdateReviewData struct {
	Attributes RevAttributes `json:"attributes"`
}

type UpdateReviewRequest struct {
	Data     UpdateReviewData `json:"data"`
	ReviewID int64
}

func NewUpdateReviewRequest(r *http.Request) (UpdateReviewRequest, error) {
	var request UpdateReviewRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal update review request")
	}

	reviewIDStr := chi.URLParam(r, "review_id")
	reviewID, err := strconv.ParseInt(reviewIDStr, 10, 64)
	if err != nil {
		return request, err
	}

	request.ReviewID = reviewID

	log.Printf("Decoded update review request: %+v", request)
	return request, nil
}
