package requests

import (
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/urlval"
	"net/http"
)

type GetReviewsListRequest struct {
	pgdb.OffsetPageParams
	FilterId []string `filter:"id"`
}

func NewGetReviewsListRequest(r *http.Request) (GetReviewsListRequest, error) {
	request := GetReviewsListRequest{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	return request, nil
}
