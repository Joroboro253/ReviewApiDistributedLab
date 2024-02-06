package handlers

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"gitlab.com/distributed_lab/ape"

	"review_api/internal/service/helpers"
	"review_api/internal/service/requests"
	"review_api/resources"
)

func UpdateRating(w http.ResponseWriter, r *http.Request) {
	request, ratingId, err := requests.NewUpdateRatingRequest(r)

	if err != nil {
		logrus.WithError(err).Error("Failed to create update rating request")
		ape.RenderErr(w, helpers.NewInvalidParamsError())
		return
	}

	ratingQ := helpers.RatingsQ(r)
	var updateData resources.UpdateRatingData

	if request.Data.Attributes.UserId != 0 {
		updateData.Attributes.UserId = request.Data.Attributes.UserId
	}
	if request.Data.Attributes.Rating != 0 {
		updateData.Attributes.Rating = request.Data.Attributes.Rating
	}

	err = ratingQ.UpdateRating(ratingId, updateData)
	if err != nil {
		logrus.WithError(err).Error("Failed to update rating")
		ape.RenderErr(w, helpers.NewInternalServerError())
		return
	}

	w.WriteHeader(http.StatusOK)
}
