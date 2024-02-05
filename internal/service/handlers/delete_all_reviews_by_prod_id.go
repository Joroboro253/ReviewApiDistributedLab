package handlers

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"gitlab.com/distributed_lab/ape"

	"review_api/internal/service/helpers"
	"review_api/internal/service/requests"
)

func DeleteReviews(w http.ResponseWriter, r *http.Request) {
	request, err := requests.DeleteReviewRequestByProductID(r)
	if err != nil {
		logrus.WithError(err).Error("Failed to create delete all review request")
		ape.RenderErr(w, helpers.NewInvalidParamsError())
		return
	}

	err = helpers.ReviewsQ(r).DeleteAllByProductId(request.ProductId)
	if err != nil {
		logrus.WithError(err).Error("Failed to delete review")
		ape.RenderErr(w, helpers.NewInternalServerError())
		return
	}

	w.WriteHeader(http.StatusOK)
}
