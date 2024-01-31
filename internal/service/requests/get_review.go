package requests

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/logan/v3"

	"review_api/internal/service/helpers"
	"review_api/resources"
)

func NewGetReviewRequest(r *http.Request) (resources.GetReviewRequest, error) {
	request := resources.GetReviewRequest{}
	request.ReviewId = cast.ToInt64(chi.URLParam(r, "id"))

	helpers.Log(r).WithFields(logan.F{"query_params": r.URL.Query()}).Info("Query Parameters")

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

	return request, nil
}
