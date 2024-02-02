package requests

import (
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"

	"review_api/resources"
)

func NewCreateReviewRequest(r *http.Request) (resources.CreateReviewRequest, error) {
	var request resources.CreateReviewRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal create review request")
	}
	if err := ValidateCreateReviewRequest(request); err != nil {
		return request, errors.Wrap(err, "Validation failed")
	}
	return request, nil
}

func ValidateCreateReviewRequest(r resources.CreateReviewRequest) error {
	return validation.Errors{
		"/data":                    validation.Validate(&r.Data, validation.Required),
		"/data/type":               validation.Validate(&r.Data.Type, validation.Required, validation.In("review")),
		"/data/id":                 validation.Validate(&r.Data.Id, validation.Required, validation.Min(0)),
		"/data/attributes":         validation.Validate(&r.Data.Attributes, validation.Required),
		"/data/attributes/content": validation.Validate(&r.Data.Attributes.Content, validation.Required, validation.Length(10, 255)),
		"/data/attributes/userId":  validation.Validate(&r.Data.Attributes.UserId, validation.Required, validation.Min(1)),
	}.Filter()
}
