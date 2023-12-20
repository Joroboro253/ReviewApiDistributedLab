package handlers

import (
	"github.com/jmoiron/sqlx"
)

type UpdateReview struct {
	DB *sqlx.DB
}

//func (h *UpdateReview) UpdateReviewById(w http.ResponseWriter, r *http.Request) *resources.APIError {
//	productId, err := strconv.Atoi(chi.URLParam(r, "product_id"))
//	if err != nil {
//		return resources.NewAPIError(http.StatusBadRequest, "StatusBadRequest", "Wrong format product_id")
//	}
//	reviewID, err := strconv.Atoi(chi.URLParam(r, "review_id"))
//	if err != nil {
//		return resources.NewAPIError(http.StatusBadRequest, "StatusBadRequest", "Error during JSON decoding")
//	}
//	// Decoding
//	var req resources.ReviewUpdateRequest
//	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
//		return resources.NewAPIError(http.StatusBadRequest, "StatusBadRequest", "Error during JSON decoding")
//	}
//	var reqBody resources.UpdateRequest
//	review := reqBody.Data.Attributes
//	// Validation
//	if validationErr := helpers.ValidateReviewAttributes(&review); validationErr != nil {
//		return validationErr
//	}
//
//	updateData := req.Data.Attributes
//	validate := validator.New()
//	if err := validate.Struct(updateData); err != nil {
//		return resources.NewAPIError(http.StatusBadRequest, "StatusBadRequest", "Data validation error")
//	}
//
//	reviewService := pg.NewReviewService(h.DB)
//	updatedReviewID, err := reviewService.UpdateReview(productId, reviewID, updateData)
//
//	if err != nil {
//		if errors.Is(err, sql.ErrNoRows) {
//			return resources.NewAPIError(http.StatusNotFound, "StatusNotFound", "Object not found")
//		}
//		return resources.NewAPIError(http.StatusBadRequest, "StatusBadRequest", "Error inserting revocation into database")
//
//	}
//
//	successResp := resources.SuccessResponse{}
//	successResp.Data.Type = "review"
//	successResp.Data.ID = updatedReviewID
//	successResp.Data.Attributes = map[string]interface{}{
//		"message": fmt.Sprintf("Review with ID %d for product %d updated successfully", updatedReviewID, productId),
//	}
//
//	w.Header().Set("Content-Type", "application/vnd.api+json")
//	w.WriteHeader(http.StatusOK)
//	json.NewEncoder(w).Encode(successResp)
//	return nil
//}
//
//func (s *ReviewService) UpdateReview(productId, reviewId int, updateData resources.ReviewUpdate) (int, error) {
//	// Initialization of SQL-builder queries
//	builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
//		Update("reviews")
//
//	// Add Update Conditions if provided
//	if updateData.UserID != nil {
//		builder = builder.Set("user_id", *updateData.UserID)
//	}
//	if updateData.Content != nil {
//		builder = builder.Set("content", *updateData.Content)
//	}
//
//	query, args, err := builder.Where(squirrel.Eq{"id": reviewId, "product_id": productId}).
//		Suffix("RETURNING id").
//		ToSql()
//	if err != nil {
//		return 0, err
//	}
//
//	// Executing the query
//	var updatedReviewID int
//	err = s.DB.QueryRow(query, args...).Scan(&updatedReviewID)
//	if err != nil {
//		return 0, err
//	}
//
//	return updatedReviewID, nil
//}
