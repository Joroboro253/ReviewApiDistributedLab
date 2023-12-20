package helpers

import (
	"github.com/jmoiron/sqlx"
)

type CreateHandler struct {
	DB *sqlx.DB
}

//func (h *CreateHandler) CreateReview(productId int, reqBody resources.UpdateRequest) (resources.Review, *resources.APIError) {
//	review := reqBody.Data.Attributes
//	review.ProductID = productId
//	review.CreatedAt = time.Now()
//	review.UpdatedAt = time.Now()
//
//	if validationErr := ValidateReviewAttributes(&review); validationErr != nil {
//		return resources.Review{}, validationErr
//	}
//
//	reviewService := pg.NewReviewService(h.DB)
//	reviewID, err := reviewService.CreateReview(&review)
//	if err != nil {
//		log.Printf("Error inserting review into database: %v", err)
//		return resources.Review{}, resources.NewAPIError(http.StatusBadRequest, "StatusBadRequest", "Error inserting review into database")
//	}
//	review.ID = reviewID
//	return review, nil
//}
//
//func (h *CreateHandler) GenerateResponse(w http.ResponseWriter) *resources.APIError {
//	w.Header().Set("Content-Type", "application/vnd.api+json")
//	w.WriteHeader(http.StatusCreated)
//	return nil
//}
