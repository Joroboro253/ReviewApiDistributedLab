package requests

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/logan/v3/errors"

	"review_api/resources"
)

func NewCreateRatingRequest(r *http.Request) (resources.CreateRatingRequest, int64, error) {
	var request resources.CreateRatingRequest
	reviewId := cast.ToInt64(chi.URLParam(r, "review_id"))

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		logrus.WithError(err).Error("Failed to decode create rating request")
		return request, reviewId, errors.Wrap(err, "Failed to unmarshal create rating request")
	}
	if err := ValidateCreateRatingRequest(request); err != nil {
		logrus.WithError(err).Error("Validation of create rating request failed")
		return request, reviewId, errors.Wrap(err, "Validation failed")
	}
	return request, reviewId, nil
}

func ValidateCreateRatingRequest(r resources.CreateRatingRequest) error {
	return validation.Errors{
		"/data":                   validation.Validate(&r.Data, validation.Required),
		"/data/type":              validation.Validate(&r.Data.Type, validation.Required, validation.In("rating")),
		"/data/attributes":        validation.Validate(&r.Data.Attributes, validation.Required),
		"/data/attributes/rating": validation.Validate(&r.Data.Attributes.Rating, validation.Required, validation.Min(1), validation.Max(5)),
		"/data/attributes/userId": validation.Validate(&r.Data.Attributes.UserId, validation.Required, validation.Min(1)),
	}.Filter()
}
