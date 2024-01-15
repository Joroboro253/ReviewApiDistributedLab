package requests

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/spf13/cast"

	"gitlab.com/distributed_lab/urlval"
)

type GetReviewRequest struct {
	ReviewID       int64  `url:"-"`
	SortBy         string `url:"sort"`
	Page           int64  `url:"page"`
	Limit          int64  `url:"limit"`
	IncludeRatings bool   `url:"include_ratings"`
}

func NewGetReviewRequest(r *http.Request) (GetReviewRequest, error) {
	request := GetReviewRequest{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	if request.Page == 0 {
		request.Page = 1
	}

	if request.Limit == 0 {
		request.Limit = 10
	}

	if request.SortBy == "" {
		request.SortBy = "date"
	}

	request.ReviewID = cast.ToInt64(chi.URLParam(r, "id"))

	return request, nil
}
