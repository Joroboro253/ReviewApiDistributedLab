package requests

import (
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"

	"review_api/resources"
)

func NewUpdateReviewRequest(r *http.Request) (resources.UpdateReviewRequest, error) {
	var request resources.UpdateReviewRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal update review request")
	}
	if err := ValidateUpdateReviewRequest(request); err != nil {
		return request, errors.Wrap(err, "Validation failed")
	}
	return request, nil
}

func ValidateUpdateReviewRequest(r resources.UpdateReviewRequest) error {
	errs := validation.Errors{
		"/data":                    validation.Validate(&r.Data, validation.Required),
		"/data/type":               validation.Validate(&r.Data.Type, validation.Required, validation.In("review")),
		"/data/reviewId":           validation.Validate(&r.Data.Id, validation.Required, validation.Min(0)),
		"/data/attributes":         validation.Validate(&r.Data.Attributes, validation.Required),
		"/data/attributes/content": validation.Validate(&r.Data.Attributes.Content, validation.Length(10, 255)),
		"/data/attributes/userId":  validation.Validate(&r.Data.Attributes.UserId, validation.Min(1)),
	}
	if r.Data.Attributes.Content == "" && r.Data.Attributes.UserId == 0 {
		errs["/data/attributes/update"] = errors.New("At least one update field must be provided")
	}
	return errs.Filter()
}
