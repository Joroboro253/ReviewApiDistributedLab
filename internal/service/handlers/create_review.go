package handlers

import (
	"net/http"
	"review_api/internal/data"
	"review_api/internal/service/helpers"
	"review_api/internal/service/requests"
)

func CreateReview(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateReviewRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("Wrong request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reviewQ, ok := helpers.ReviewsQ(r).(data.ReviewQ)
	if !ok || reviewQ == nil {
		helpers.Log(r).WithError(err).Error("ReviewQ is not available in the request context")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = reviewQ.Transaction(func(q data.ReviewQ) error {
		review := data.Review{
			ProductID: request.Data.ProductID,
			UserID:    request.Data.UserID,
			Content:   request.Data.Content,
		}

		_, err = q.Insert(review)
		return nil
	})

	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to create review")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
