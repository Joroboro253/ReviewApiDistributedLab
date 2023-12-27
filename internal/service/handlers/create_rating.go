package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
	"review_api/internal/data"
	"review_api/internal/service/helpers"
	"review_api/internal/service/requests"
)

func CreateRating(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateRatingRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("Wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	ratingQ := helpers.RatingsQ(r)
	if ratingQ == nil {
		helpers.Log(r).Error("RatingQ is not available in the request context")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	resultRating, err := ratingQ.Insert(data.Rating{
		ReviewID: request.Data.ReviewID,
		UserID:   request.Data.UserID,
		Rating:   request.Data.Rating,
	})
	if err != nil {
		helpers.Log(r).WithError(err).Error("Failed to create rating")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ratingModel := data.Rating{
		ID:        resultRating.ID,
		ReviewID:  resultRating.ReviewID,
		UserID:    resultRating.UserID,
		Rating:    resultRating.Rating,
		CreatedAt: resultRating.CreatedAt,
		UpdateAt:  resultRating.UpdateAt,
	}

	result := data.RatingResponse{
		Data: ratingModel,
	}
	ape.Render(w, result)
}
