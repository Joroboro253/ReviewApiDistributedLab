package handlers

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"gitlab.com/distributed_lab/ape"

	"review_api/internal/service/helpers"
	"review_api/internal/service/requests"
)

func DeleteRating(w http.ResponseWriter, r *http.Request) {
	request, err := requests.DeleteRatingRequest(r)
	if err != nil {
		logrus.WithError(err).Error("Failed to create delete rating request")
		ape.RenderErr(w, helpers.NewInvalidParamsError())
		return
	}

	err = helpers.RatingsQ(r).DeleteRating(request.RatingId)
	if err != nil {
		logrus.WithError(err).Error("Failed to delete rating")
		ape.RenderErr(w, helpers.NewInternalServerError())
		return
	}

	w.WriteHeader(http.StatusOK)
}
