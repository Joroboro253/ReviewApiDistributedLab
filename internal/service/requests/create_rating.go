package requests

import (
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"

	"review_api/resources"
)

func NewCreateRatingRequest(r *http.Request) (resources.CreateRatingRequest, error) {
	var request resources.CreateRatingRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "Failed to unmarshal create rating request")
	}
	if err := ValidateCreateRatingRequest(request); err != nil {
		return request, errors.Wrap(err, "Validation failed")
	}
	return request, nil
}

func ValidateCreateRatingRequest(r resources.CreateRatingRequest) error {
	return validation.Errors{
		"/data":                     validation.Validate(&r.Data, validation.Required),
		"/data/type":                validation.Validate(&r.Data.Type, validation.Required, validation.In("rating")),
		"/data/ratingId":            validation.Validate(&r.Data.RatingId, validation.Required, validation.Min(1)),
		"/data/attributes":          validation.Validate(&r.Data.Attributes, validation.Required),
		"/data/attributes/rating":   validation.Validate(&r.Data.Attributes.Rating, validation.Required, validation.Min(float64(1)), validation.Max(float64(5))),
		"/data/attributes/reviewId": validation.Validate(&r.Data.Attributes.ReviewId, validation.Required, validation.Min(1)),
		"/data/attributes/userId":   validation.Validate(&r.Data.Attributes.UserId, validation.Required, validation.Min(1)),
	}.Filter()
}
