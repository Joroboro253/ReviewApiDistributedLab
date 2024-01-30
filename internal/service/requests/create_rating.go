package requests

import (
	"encoding/json"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3/errors"

	"review_api/resources"
)

func NewCreateRatingRequest(r *http.Request) (resources.CreateRatingRequest, error) {
	var request resources.CreateRatingRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "Failed to unmarshal")
	}

	return request, nil
}

//func (r *CreateRatingRequest) Validate() error {
//	return validation.Errors{
//		"data/attributes/content": validation.Validate(&r.Data.Attributes.Rating, validation.Required, validation.Min(1), validation.Max(5)),
//	}.Filter()
//}
