package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

	"review_api/internal/service/helpers"
	"review_api/internal/service/requests"
	"review_api/resources"
)

func UpdateReview(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewUpdateReviewRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	reviewQ := helpers.ReviewsQ(r)
	var updateData resources.UpdateReviewData

	if request.Data.Attributes.ProductID != nil {
		updateData.ProductId = request.Data.Attributes.ProductID
	}
	if request.Data.Attributes.UserID != nil {
		updateData.UserId = request.Data.Attributes.UserID
	}
	if request.Data.Attributes.Content != nil {
		updateData.Content = request.Data.Attributes.Content
	}

	_, err = reviewQ.UpdateReview(request.ReviewID, updateData)
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusCreated)
}
