package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"

	"review_api/internal/service/helpers"
	"review_api/internal/service/requests"
	"review_api/resources"
)

func UpdateRating(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewUpdateRatingRequest(r)
	if err != nil {
		ape.RenderErr(w, helpers.NewInvalidParamsError())
		return
	}

	ratingQ := helpers.RatingsQ(r)
	var updateData resources.UpdateRatingData

	if request.Data.Id != 0 {
		updateData.Id = request.Data.Id
	}
	if request.Data.Attributes.UserId != 0 {
		updateData.Attributes.UserId = request.Data.Attributes.UserId
	}
	if request.Data.Attributes.Rating != 0 {
		updateData.Attributes.Rating = request.Data.Attributes.Rating
	}

	_, err = ratingQ.UpdateRating(request.Data.Id, updateData)
	if err != nil {
		ape.RenderErr(w, helpers.NewInternalServerError())
		return
	}

	w.WriteHeader(http.StatusCreated)
}
