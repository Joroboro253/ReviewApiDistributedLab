package handlers

import (
	"net/http"
	"review_api/internal/data"
	"review_api/internal/service/helpers"
	"review_api/internal/service/requests"
)

func CreateRating(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateRatingRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("Wrong request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ratingQ := helpers.RatingsQ(r)
	if ratingQ == nil {
		helpers.Log(r).Error("RatingQ is not available in the request context")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = ratingQ.Insert(data.Rating{
		ReviewID: request.Data.ReviewID,
		UserID:   request.Data.UserID,
		Rating:   request.Data.Rating,
	})
	if err != nil {
		helpers.Log(r).WithError(err).Error("Failed to create rating")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
