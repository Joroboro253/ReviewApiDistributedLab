package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

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

	if request.Data.Attributes.ReviewID != nil {
		updateData.ReviewId = request.Data.Attributes.ReviewID
	}
	if request.Data.Attributes.UserID != nil {
		updateData.UserId = request.Data.Attributes.UserID
	}
	if request.Data.Attributes.Rating != nil {
		updateData.Rating = request.Data.Attributes.Rating
	}

	_, err = ratingQ.UpdateRating(request.RatingID, updateData)
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusCreated)
}
