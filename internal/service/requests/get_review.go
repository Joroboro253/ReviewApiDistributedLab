package requests

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/spf13/cast"

	"review_api/resources"
)

func NewGetReviewRequest(r *http.Request) (resources.ReviewQueryParams, error) {
	request := resources.ReviewQueryParams{}
	request.ProductId = cast.ToInt64(chi.URLParam(r, "product_id"))

	includeRatingsParam := r.URL.Query().Get("includeRatings")
	if includeRatingsParam != "" {
		request.IncludeRatings, _ = strconv.ParseBool(includeRatingsParam)
	}

	limitParam := r.URL.Query().Get("limit")
	if limitParam != "" {
		request.Limit, _ = strconv.ParseInt(limitParam, 10, 64)
	} else {
		request.Limit = 10
	}

	pageParam := r.URL.Query().Get("page")
	if pageParam != "" {
		request.Page, _ = strconv.ParseInt(pageParam, 10, 64)
	} else {
		request.Page = 1
	}

	sortByParam := r.URL.Query().Get("sortBy")
	if sortByParam != "" {
		request.SortBy = sortByParam
	} else {
		request.SortBy = "date"
	}

	sortDirectionParam := r.URL.Query().Get("sortDirection")
	if sortDirectionParam != "" {
		request.SortDirection = sortDirectionParam
	} else {
		request.SortDirection = "asc"
	}

	return request, nil
}
