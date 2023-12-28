package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
	"review_api/internal/service/helpers"
	"review_api/internal/service/requests"
)

func DeleteRating(w http.ResponseWriter, r *http.Request) {
	request, err := requests.DeleteRatingRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	err = helpers.RatingsQ(r).DeleteRating(request.RatingID)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed delete review from DB")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Rating deleted successfully"))
}
