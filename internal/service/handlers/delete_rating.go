package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"

	"review_api/internal/service/helpers"
	"review_api/internal/service/requests"
)

func DeleteRating(w http.ResponseWriter, r *http.Request) {
	request, err := requests.DeleteRatingRequest(r)
	if err != nil {
		ape.RenderErr(w, helpers.NewInvalidParamsError())
		return
	}

	err = helpers.RatingsQ(r).DeleteRating(request.RatingId)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed delete review from DB")
		ape.RenderErr(w, helpers.NewInternalServerError())
		return
	}

	w.WriteHeader(http.StatusOK)
}
