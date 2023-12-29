package requests

import (
	"encoding/json"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"log"
	"net/http"
	"review_api/internal/data"
)

type CreateReviewRequest struct {
	Data data.Review `json:"data"`
}

func NewCreateReviewRequest(r *http.Request) (CreateReviewRequest, error) {
	var request CreateReviewRequest
	log.Printf("NewCreateReviewRequests")

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}

	return request, nil
}
