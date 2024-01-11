package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

	"review_api/internal/data"
	"review_api/internal/service/helpers"
	"review_api/internal/service/requests"
)

func CreateRating(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateRatingRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	err = helpers.RatingsQ(r).Insert(data.Rating{
		ReviewID: request.Data.ReviewID,
		UserID:   request.Data.UserID,
		Rating:   request.Data.Rating,
	})
	if err != nil {
		helpers.Log(r).WithError(err).Error("Failed to create rating")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusCreated)
}
