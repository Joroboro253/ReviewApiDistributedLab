package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

	"review_api/internal/service/helpers"
	"review_api/internal/service/requests"
	"review_api/resources"
)

func GetReviews(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetReviewRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	productId, err := strconv.ParseInt(chi.URLParam(r, "product_id"), 10, 64)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	reviewQ := helpers.ReviewsQ(r)
	sortParam := resources.SortParam{Limit: request.Limit, Page: request.Page, SortBy: request.SortBy, SortDirection: request.SortDirection}

	helpers.Log(r).WithError(err).Debugf("Sorting params in handler: %+v", sortParam)

	reviews, meta, err := reviewQ.Select(sortParam, request.IncludeRatings, productId)
	if err != nil {
		helpers.Log(r).WithError(err).Info("Internal server Error")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, helpers.ConvertToAPIResponse(reviews, meta))
}
