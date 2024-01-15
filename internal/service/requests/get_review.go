package requests

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/logan/v3"

	"gitlab.com/distributed_lab/urlval"

	"review_api/internal/service/helpers"
)

type GetReviewRequest struct {
	ReviewID       int64  `url:"-"`
	SortBy         string `url:"sort"`
	Page           int64  `url:"page"`
	Limit          int64  `url:"limit"`
	IncludeRatings bool   `url:"include_ratings"`
	SortDirection  string `url:"sort_dir"`
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

	if request.SortDirection == "" {
		request.SortDirection = "asc"
	}

	request.ReviewID = cast.ToInt64(chi.URLParam(r, "id"))
	helpers.Log(r).WithFields(logan.F{"request": request}).Info("Parsed GetReviewRequest")

	return request, nil
}
