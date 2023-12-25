package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
	"review_api/internal/service/helpers"
	"review_api/internal/service/requests"
)

func DeleteReviewByID(w http.ResponseWriter, r *http.Request) {
	request, err := requests.DeleteReviewRequestByReviewID(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	err = helpers.ReviewsQ(r).DeleteByReviewId(request.ReviewID)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed delete review from DB")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Review deleted successfully"))
}
