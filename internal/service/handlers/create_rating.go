package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"

	"review_api/internal/data"
	"review_api/internal/service/helpers"
	"review_api/internal/service/requests"
)

func CreateRating(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateRatingRequest(r)
	if err != nil {
		ape.RenderErr(w, helpers.NewInvalidParamsError())
		return
	}

	err = helpers.RatingsQ(r).Insert(data.Rating{
		ReviewID: request.Data.Attributes.ReviewId,
		UserID:   request.Data.Attributes.UserId,
		Rating:   request.Data.Attributes.Rating,
	})

	if err != nil {
		helpers.Log(r).WithError(err).Error("Failed to create rating")
		ape.RenderErr(w, helpers.NewInternalServerError())
		return
	}

	w.WriteHeader(http.StatusCreated)
}
