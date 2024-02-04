package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"

	"review_api/internal/data"
	"review_api/internal/service/helpers"
	"review_api/internal/service/requests"
)

func CreateReview(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateReviewRequest(r)
	if err != nil {
		ape.RenderErr(w, helpers.NewInvalidParamsError())
		return
	}

	err = helpers.ReviewsQ(r).Insert(data.Review{
		ID:        request.Data.Id,
		UserID:    request.Data.Attributes.UserId,
		Content:   request.Data.Attributes.Content,
		Rating:    request.Data.Attributes.Rating,
		ProductID: request.Data.Attributes.ProductId,
	})

	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to create review")
		ape.RenderErr(w, helpers.NewInternalServerError())
		return
	}

	w.WriteHeader(http.StatusCreated)
}
