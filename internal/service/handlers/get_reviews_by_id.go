package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
	"review_api/internal/data"
	"review_api/internal/service/helpers"
	"review_api/internal/service/requests"
	"review_api/resources"
)

func GetReviews(w http.ResponseWriter, r *http.Request) {

	_, err := requests.NewGetReviewsListRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	reviewsQ := helpers.ReviewsQ(r)
	reviews, err := reviewsQ.Select()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get reviews")
		ape.Render(w, problems.InternalError())
		return
	}

	response := resources.ReviewListResponse{
		Data: newReviewsList(reviews),
	}

	ape.Render(w, response)

}

func newReviewsList(reviews []data.Review) []data.Review {
	result := make([]data.Review, len(reviews))
	for i, review := range reviews {
		result[i] = data.Review{
			ID:        review.ID,
			ProductID: review.ProductID,
			UserID:    review.UserID,
			Content:   review.Content,
			Rating:    review.Rating,
		}
	}
	return result
}
