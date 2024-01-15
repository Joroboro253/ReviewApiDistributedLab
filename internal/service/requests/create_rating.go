package requests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

type CreateRatingRequest struct {
	Data struct {
		Type       string `json:"type"`
		Attributes struct {
			ReviewID int64   `json:"review_id"`
			UserID   int64   `json:"user_id"`
			Rating   float64 `json:"rating"`
		} `json:"attributes"`
	} `json:"data"`
}

func NewCreateRatingRequest(r *http.Request) (CreateRatingRequest, error) {
	var request CreateRatingRequest

	body, _ := ioutil.ReadAll(r.Body)

	if err := json.NewDecoder(bytes.NewBuffer(body)).Decode(&request); err != nil {
		return request, errors.Wrap(err, "Failed to unmarshal")
	}

	return request, nil
}
