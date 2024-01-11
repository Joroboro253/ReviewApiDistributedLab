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

func UpdateRating(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewUpdateRatingRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	ratingQ := helpers.RatingsQ(r)
	var updateData resources.UpdateRatingData

	if request.Data.ReviewID != nil {
		updateData.ReviewId = request.Data.ReviewID
	}
	if request.Data.UserID != nil {
		updateData.UserId = request.Data.UserID
	}
	if request.Data.Rating != nil {
		updateData.Rating = request.Data.Rating
	}

	updatedRating, err := ratingQ.UpdateRating(request.RatingID, updateData)
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}

	response := data.RatingResponse{
		Data: updatedRating,
	}
	ape.Render(w, response)
}
