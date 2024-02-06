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

func NewCreateReviewRequest(r *http.Request) (resources.CreateReviewRequest, int64, error) {
	var request resources.CreateReviewRequest
	productId := cast.ToInt64(chi.URLParam(r, "product_id"))

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		logrus.WithError(err).Error("Failed to decode create review request")
		return request, productId, errors.Wrap(err, "failed to unmarshal create review request")
	}
	if err := ValidateCreateReviewRequest(request); err != nil {
		logrus.WithError(err).Error("Validation of create review request failed")
		return request, productId, errors.Wrap(err, "Validation failed")
	}
	return request, productId, nil
}

func ValidateCreateReviewRequest(r resources.CreateReviewRequest) error {
	return validation.Errors{
		"/data":                    validation.Validate(&r.Data, validation.Required),
		"/data/type":               validation.Validate(&r.Data.Type, validation.Required, validation.In("review")),
		"/data/attributes":         validation.Validate(&r.Data.Attributes, validation.Required),
		"/data/attributes/rating":  validation.Validate(&r.Data.Attributes.Rating, validation.Required, validation.Min(1), validation.Max(5)),
		"/data/attributes/content": validation.Validate(&r.Data.Attributes.Content, validation.Required, validation.Length(10, 255)),
		"/data/attributes/userId":  validation.Validate(&r.Data.Attributes.UserId, validation.Required, validation.Min(0)),
	}.Filter()
}
