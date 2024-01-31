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

	var updateData resources.UpdateReviewData

	if request.Data.Attributes.ProductId != 0 {
		updateData.Attributes.ProductId = request.Data.Attributes.ProductId
	}
	if request.Data.Attributes.UserId != 0 {
		updateData.Attributes.UserId = request.Data.Attributes.UserId
	}
	if request.Data.Attributes.Content != "" {
		updateData.Attributes.Content = request.Data.Attributes.Content
	}

	// TODO это было перенесено из pg 20:00
	//if !updateFields {
	//	return data.Review{}, errors.New("no fields to update")
	//}

	reviewQ := helpers.ReviewsQ(r)

	_, err = reviewQ.UpdateReview(request.Data.ReviewId, updateData)
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusCreated)
}
