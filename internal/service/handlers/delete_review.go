package handlers

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"gitlab.com/distributed_lab/ape"

	"review_api/internal/service/helpers"
	"review_api/internal/service/requests"
)

func DeleteReview(w http.ResponseWriter, r *http.Request) {
	request, err := requests.DeleteReviewRequest(r)
	if err != nil {
		logrus.WithError(err).Error("Failed to create delete review request")
		ape.RenderErr(w, helpers.NewInvalidParamsError())
		return
	}
	logrus.Printf("Review id: %d", request.ReviewId)
	err = helpers.ReviewsQ(r).DeleteReview(request.ReviewId)
	if err != nil {
		logrus.WithError(err).Error("failed to delete rating")
		ape.RenderErr(w, helpers.NewInternalServerError())
		return
	}

	w.WriteHeader(http.StatusOK)
}
