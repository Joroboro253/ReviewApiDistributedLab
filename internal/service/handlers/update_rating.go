package handlers

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
	"review_api/internal/data"
	"review_api/internal/service/helpers"
	"review_api/internal/service/requests"
	"strconv"
)

func UpdateRating(w http.ResponseWriter, r *http.Request) {
	ratingIDStr := chi.URLParam(r, "rating_id")
	ratingID, err := strconv.ParseInt(ratingIDStr, 10, 64)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	request, err := requests.NewUpdateRatingRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	ratingQ := helpers.RatingsQ(r)
	updateData := make(map[string]interface{})
	if request.Data.ReviewID != 0 {
		updateData["review_id"] = request.Data.ReviewID
	}
	if request.Data.UserID != 0 {
		updateData["user_id"] = request.Data.UserID
	}
	if request.Data.Rating != 0 {
		updateData["rating"] = request.Data.Rating
	}

	updatedRating, err := ratingQ.UpdateRating(ratingID, updateData)
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}

	response := data.RatingResponse{
		Data: updatedRating,
	}
	ape.Render(w, response)
}
