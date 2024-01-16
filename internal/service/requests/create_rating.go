package requests

import (
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type CreateRatingRequest struct {
	Data struct {
		Type       string `json:"type"`
		Attributes struct {
			ReviewID int64   `json:"review_id"`
			UserID   int64   `json:"user_id"`
			Rating   float64 `json:"rating"`
		} `json:"attributes"`
	} `json:"data"`
}

func NewCreateRatingRequest(r *http.Request) (CreateRatingRequest, error) {
	var request CreateRatingRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "Failed to unmarshal")
	}

	return request, nil
}

func (r *CreateRatingRequest) Validate() error {
	return validation.Errors{
		"data/attributes/content": validation.Validate(&r.Data.Attributes.Rating, validation.Required, validation.Min(1), validation.Max(5)),
	}.Filter()
}
