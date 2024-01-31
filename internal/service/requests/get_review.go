package requests

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/spf13/cast"

	"review_api/resources"
)

func NewGetReviewRequest(r *http.Request) (resources.GetReviewRequest, error) {
	request := resources.GetReviewRequest{}
	request.ReviewId = cast.ToInt64(chi.URLParam(r, "id"))

	includeRatingsParam := r.URL.Query().Get("includeRatings")
	if includeRatingsParam != "" {
		request.IncludeRatings, _ = strconv.ParseBool(includeRatingsParam)
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
	log.Printf("Before includeRatings check: includeRatings = %v", request)

	return request, nil
}
