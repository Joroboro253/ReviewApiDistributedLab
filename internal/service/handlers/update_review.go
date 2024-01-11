package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

	"review_api/internal/data"
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

	if request.Data.ProductID != nil {
		updateData.ProductId = request.Data.ProductID
	}
	if request.Data.UserID != nil {
		updateData.UserId = request.Data.UserID
	}
	if request.Data.Content != nil {
		updateData.Content = request.Data.Content
	}

	updatedReview, err := reviewQ.UpdateReview(request.ReviewID, updateData)
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}

	response := data.ReviewResponse{
		Data: updatedReview,
	}
	ape.Render(w, response)
}
