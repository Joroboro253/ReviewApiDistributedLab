package handlers

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"gitlab.com/distributed_lab/ape"

	"review_api/internal/service/helpers"
	"review_api/internal/service/requests"
	"review_api/resources"
)

func GetReviews(w http.ResponseWriter, r *http.Request) {
	request, productId, err := requests.NewGetReviewRequest(r)
	if err != nil {
		logrus.WithError(err).Error("Failed to create get review request")
		ape.RenderErr(w, helpers.NewInvalidParamsError())
		return
	}

	reviewQ := helpers.ReviewsQ(r)
	sortParam := resources.SortParam{Limit: request.Limit, Page: request.Page, SortBy: request.SortBy, SortDirection: request.SortDirection}

	reviews, meta, err := reviewQ.Select(sortParam, request.IncludeRatings, productId)
	if err != nil {
		logrus.WithError(err).Error("Failed to get review")
		ape.RenderErr(w, helpers.NewInternalServerError())
		return
	}

	ape.Render(w, helpers.ConvertToJSONResponse(reviews, meta))
}
