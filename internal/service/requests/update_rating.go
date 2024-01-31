package requests

import (
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"

	"review_api/resources"
)

func NewUpdateRatingRequest(r *http.Request) (resources.UpdateRatingRequest, error) {
	var request resources.UpdateRatingRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal update rating request")
	}
	if err := ValidateUpdateRatingRequest(request); err != nil {
		return request, errors.Wrap(err, "Validation failed")
	}
	return request, nil
}

func ValidateUpdateRatingRequest(r resources.UpdateRatingRequest) error {
	errs := validation.Errors{
		"/data":                     validation.Validate(&r.Data, validation.Required),
		"/data/type":                validation.Validate(&r.Data.Type, validation.Required, validation.In("rating")),
		"/data/ratingId":            validation.Validate(&r.Data.RatingId, validation.Required, validation.Min(1)),
		"/data/attributes":          validation.Validate(&r.Data.Attributes, validation.Required),
		"/data/attributes/rating":   validation.Validate(&r.Data.Attributes.Rating, validation.Min(float64(1)), validation.Max(float64(5))),
		"/data/attributes/reviewId": validation.Validate(&r.Data.Attributes.ReviewId, validation.Min(1)),
		"/data/attributes/userId":   validation.Validate(&r.Data.Attributes.UserId, validation.Min(1)),
	}
	if r.Data.Attributes.ReviewId == 0 && r.Data.Attributes.UserId == 0 && r.Data.Attributes.Rating == 0 {
		errs["data/attributes/update"] = errors.New("At least one update fields must be provided")
	}
	return errs.Filter()
}
