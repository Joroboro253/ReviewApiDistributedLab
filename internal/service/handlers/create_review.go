package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"log"
	"net/http"
	"review_api/internal/data/pg"
	"review_api/internal/service/helpers"
	"review_api/internal/service/requests"
	"review_api/resources"
)

func CreateReview(w http.ResponseWriter, r *http.Request) {
	db, _ := helpers.GetDBFromContext(r) // Предполагаем, что функция GetDBFromContext определена в helpers

	logger := log.Default() // Использование стандартного логгера

	productId, apiErr := helpers.GetProductIDFromURL(r)
	if apiErr != nil {
		ape.RenderErr(w, problems.BadRequest(apiErr)...)
		logger.Println("Error getting product ID from URL:", apiErr)
		return
	}

	reqBody, apiErr := requests.DecodeReviewRequestBody(r)
	if apiErr != nil {
		ape.RenderErr(w, problems.BadRequest(apiErr)...)
		logger.Printf("Error decoding review request body: %v\n", apiErr)
		return
	}

	if reqBody.Data.Type != "review" {
		apiErr := resources.NewAPIError(http.StatusBadRequest, "StatusBadRequest", "Incorrect data type")
		ape.RenderErr(w, problems.BadRequest(apiErr)...)
		logger.Printf("Incorrect data type: %v\n", apiErr)
		return
	}

	reviewService := pg.NewReviewService(db)
	review, apiErr := reviewService.CreateReview(productId, reqBody)
	if apiErr != nil {
		ape.RenderErr(w, problems.InternalError())
		logger.Printf("Failed to create review: %v\n", apiErr)
		return
	}

	helpers.GenerateReviewResponse(w, review)
}
