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

func NewUpdateReviewRequest(r *http.Request) (resources.UpdateReviewRequest, int64, error) {
	var request resources.UpdateReviewRequest
	reviewId := cast.ToInt64(chi.URLParam(r, "review_id"))

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		logrus.WithError(err).Error("Failed to unmarshal update review request")
		return request, reviewId, errors.Wrap(err, "failed to unmarshal update review request")
	}
	if err := ValidateUpdateReviewRequest(request); err != nil {
		logrus.WithError(err).Error("Validation update review request failed")
		return request, reviewId, errors.Wrap(err, "Validation failed")
	}
	return request, reviewId, nil
}

func ValidateUpdateReviewRequest(r resources.UpdateReviewRequest) error {
	errs := validation.Errors{
		"/data":                    validation.Validate(&r.Data, validation.Required),
		"/data/type":               validation.Validate(&r.Data.Type, validation.Required, validation.In("review")),
		"/data/attributes":         validation.Validate(&r.Data.Attributes, validation.Required),
		"/data/attributes/content": validation.Validate(&r.Data.Attributes.Content, validation.Length(10, 255)),
		"/data/attributes/userId":  validation.Validate(&r.Data.Attributes.UserId, validation.Min(1)),
		"/data/attributes/rating":  validation.Validate(&r.Data.Attributes.Rating, validation.Min(1), validation.Max(5)),
	}
	if r.Data.Attributes.Content == "" && r.Data.Attributes.UserId == 0 && r.Data.Attributes.Rating == 0 {
		errs["/data/attributes/update"] = errors.New("At least one update field must be provided")
	}
	return errs.Filter()
}
