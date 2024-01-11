package requests

import (
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type CreateReviewRequest struct {
	Data struct {
		Type       string `json:"type"`
		Attributes struct {
			ProductID int64  `json:"product_id"`
			UserID    int64  `json:"user_id"`
			Content   string `json:"content"`
		} `json:"attributes"`
	} `json:"data"`
}

func NewCreateReviewRequest(r *http.Request) (CreateReviewRequest, error) {
	var request CreateReviewRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}
	if err := request.Validate(); err != nil {
		return request, errors.Wrap(err, "validation failed")
	}

	return request, nil
}

func (r *CreateReviewRequest) Validate() error {
	return validation.Errors{
		"/data/attributes/content": validation.Validate(&r.Data.Attributes.Content, validation.Required, validation.Length(10, 255)),
	}.Filter()
}
