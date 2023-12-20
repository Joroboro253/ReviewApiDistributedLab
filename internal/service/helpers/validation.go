package helpers

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"review_api/resources"
)

var validate = validator.New()

func ValidateReviewAttributes(attributes *resources.Review) *resources.APIError {
	if err := validate.Struct(attributes); err != nil {
		return resources.NewAPIError(http.StatusBadRequest, "StatusBadRequest", "Error during JSON decoding")
	}
	return nil
}
