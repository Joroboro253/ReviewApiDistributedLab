package handlers

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"gitlab.com/distributed_lab/ape"

	"review_api/internal/service/helpers"
	"review_api/internal/service/requests"
	"review_api/resources"
)

func UpdateReview(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewUpdateReviewRequest(r)
	if err != nil {
		logrus.WithError(err).Error("Failed to create update review request")
		ape.RenderErr(w, helpers.NewInvalidParamsError())
		return
	}

	var updateData resources.UpdateReviewData
	updateData.Id = request.Data.Id
	if request.Data.Attributes.ProductId != 0 {
		updateData.Attributes.ProductId = request.Data.Attributes.ProductId
	}
	if request.Data.Attributes.UserId != 0 {
		updateData.Attributes.UserId = request.Data.Attributes.UserId
	}
	if request.Data.Attributes.Content != "" {
		updateData.Attributes.Content = request.Data.Attributes.Content
	}

	err = helpers.ReviewsQ(r).UpdateReview(updateData)
	if err != nil {
		logrus.WithError(err).Error("Failed to update review")
		ape.RenderErr(w, helpers.NewInternalServerError())
		return
	}

	w.WriteHeader(http.StatusOK)
}
