package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"

	"review_api/internal/service/helpers"
	"review_api/internal/service/requests"
)

func DeleteReviews(w http.ResponseWriter, r *http.Request) {
	request, err := requests.DeleteReviewRequestByProductID(r)
	if err != nil {
		ape.RenderErr(w, helpers.NewInvalidParamsError())
		return
	}

	err = helpers.ReviewsQ(r).DeleteAllByProductId(request.ProductId)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to delete all reviews by product id from DB")
		ape.RenderErr(w, helpers.NewInternalServerError())
		return
	}

	w.WriteHeader(http.StatusOK)
}
