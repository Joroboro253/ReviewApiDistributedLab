package requests

import (
	"encoding/json"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
)

type UpdateReviewRequest struct {
	Data struct {
		ProductID int     `json:"product_id"`
		UserID    int     `json:"user_id"`
		Content   string  `json:"content"`
		Rating    float32 `json:"rating"`
	} `json:"data"`
}

func NewUpdateReviewRequest(r *http.Request) (UpdateReviewRequest, error) {
	var request UpdateReviewRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal update review request")
	}
	return request, nil
}
