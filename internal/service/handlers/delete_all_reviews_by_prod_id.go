package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
	"review_api/internal/service/helpers"
	"review_api/internal/service/requests"
)

func DeleteAllByProductId(w http.ResponseWriter, r *http.Request) {
	// Assuming NewDeleteReviewRequest parses the request and extracts the review ID
	request, err := requests.DeleteReviewRequestByProductID(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	_, err = helpers.ReviewsQ(r).FilterByID(request.ProductID).Get()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get product from DB")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	err = helpers.ReviewsQ(r).DeleteAllByProductId(request.ProductID)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to delete all reviews by product id from DB")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Review deleted successfully"))
}
