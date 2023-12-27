package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"log"
	"net/http"
	"review_api/internal/data"
	"review_api/internal/service/helpers"
	"review_api/internal/service/requests"
)

const createReview = "createReview"

func CreateReview(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateReviewRequest(r)
	if err != nil {
		log.Print("Wrong request")
		helpers.Log(r).WithError(err).Info("Wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	reviewQ, ok := helpers.ReviewsQ(r).(data.ReviewQ)
	if !ok || reviewQ == nil {
		helpers.Log(r).WithError(err).Error("ReviewQ is not available in the request context")
		errors.Wrap(err, "failed to insert review (helpers.ReviewsQ(r))")
		return
	}

	var resultReview data.Review
	log.Print("Transaction")
	err = helpers.ReviewsQ(r).Transaction(func(q data.ReviewQ) error {
		review := data.Review{
			ProductID: request.Data.ProductID,
			UserID:    request.Data.UserID,
			Content:   request.Data.Content,
		}

		resultReview, err = q.Insert(review)
		if err != nil {
			return errors.Wrap(err, "failed to insert review")
		}

		return nil
	})

	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to create review")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	reviewModel := data.Review{
		ID:        resultReview.ID,
		ProductID: resultReview.ProductID,
		UserID:    resultReview.UserID,
		Content:   resultReview.Content,
		CreatedAt: resultReview.CreatedAt,
		UpdatedAt: resultReview.UpdatedAt,
	}

	result := data.ReviewResponse{
		Data: reviewModel,
	}
	ape.Render(w, result)
}
