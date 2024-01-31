package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3"

	"review_api/internal/data"
	"review_api/internal/service/helpers"
	"review_api/internal/service/requests"
	"review_api/resources"
)

func GetReviews(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetReviewRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("Wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	helpers.Log(r).WithFields(logan.F{"request": request}).Info("Received GetReviewRequest")

	reviewQ := helpers.ReviewsQ(r)
	sortParam := resources.SortParam{Limit: request.Limit, Page: request.Page, SortBy: request.SortBy, SortDirection: request.SortDirection}
	reviews, meta, err := reviewQ.Select(sortParam, request.IncludeRatings)
	if err != nil {
		helpers.Log(r).WithError(err).Info("Internal server Error")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	response := struct {
		Data []data.ReviewWithRatings `json:"data"`
		Meta *resources.PaginationMeta
	}{
		Data: reviews,
		Meta: meta,
	}
	helpers.Log(r).WithField("response", response).Info("Sending response")
	ape.Render(w, response)
}
