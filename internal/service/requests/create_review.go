package requests

import (
	"encoding/json"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3/errors"

	"review_api/resources"
)

func NewCreateReviewRequest(r *http.Request) (resources.CreateReviewRequest, error) {
	var request resources.CreateReviewRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}
	//if err := request.Validate(); err != nil {
	//	return request, errors.Wrap(err, "validation failed")
	//}

	return request, nil
}

//func (r *CreateReviewRequest) Validate() error {
//	return validation.Errors{
//		"/data/attributes/content": validation.Validate(&r.Data.Attributes.Content, validation.Required, validation.Length(10, 255)),
//	}.Filter()
//}
