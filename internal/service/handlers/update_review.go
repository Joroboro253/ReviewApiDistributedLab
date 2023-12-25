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

func UpdateReview(w http.ResponseWriter, r *http.Request) {
	reviewIDStr := chi.URLParam(r, "review_id")
	reviewID, err := strconv.ParseInt(reviewIDStr, 10, 64)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	request, err := requests.NewUpdateReviewRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	reviewQ := helpers.ReviewsQ(r)
	updateData := make(map[string]interface{})
	if request.Data.ProductID != 0 {
		updateData["product_id"] = request.Data.ProductID
	}
	if request.Data.UserID != 0 {
		updateData["user_id"] = request.Data.UserID
	}
	if request.Data.Content != "" {
		updateData["content"] = request.Data.Content
	}
	if request.Data.Rating != 0 {
		updateData["rating"] = request.Data.Rating
	}

	updatedReview, err := reviewQ.Update(reviewID, updateData)
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}

	response := data.ReviewResponse{
		Data: updatedReview,
	}
	ape.Render(w, response)
}
