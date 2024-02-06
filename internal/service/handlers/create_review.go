package handlers

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

	"review_api/internal/data"
	"review_api/internal/service/helpers"
	"review_api/internal/service/requests"
)

func CreateReview(w http.ResponseWriter, r *http.Request) {
	request, productId, err := requests.NewCreateReviewRequest(r)
	if err != nil {
		logrus.WithError(err).Error("Failed to create new review request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	err = helpers.ReviewsQ(r).Insert(data.Review{
		UserID:    request.Data.Attributes.UserId,
		Content:   request.Data.Attributes.Content,
		Rating:    request.Data.Attributes.Rating,
		ProductID: productId,
	})

	if err != nil {
		logrus.WithError(err).Error("Failed to create review")
		ape.RenderErr(w, helpers.NewInternalServerError())
		return
	}

	w.WriteHeader(http.StatusCreated)
}
