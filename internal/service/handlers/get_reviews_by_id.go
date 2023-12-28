package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
	"review_api/internal/data"
	"review_api/internal/service/helpers"
	"review_api/internal/service/requests"
)

func GetReviews(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetReviewRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	reviewQ := helpers.ReviewsQ(r)
	reviews, err := reviewQ.Select(request.SortBy, request.Page, request.Limit, request.IncludeRatings)
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}

	response := struct {
		Data []data.ReviewWithRatings `json:"data"`
	}{
		Data: reviews,
	}
	ape.Render(w, response)
}
