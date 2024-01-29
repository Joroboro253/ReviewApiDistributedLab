package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

	"review_api/internal/data"
	"review_api/internal/service/helpers"
	"review_api/internal/service/requests"
)

func CreateReview(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateReviewRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	err = helpers.ReviewsQ(r).Insert(data.Review{
		ProductID: request.Data.Attributes.ReviewId,
		UserID:    request.Data.Attributes.UserId,
		Content:   request.Data.Attributes.Content,
	})

	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to create review")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusCreated)
}
