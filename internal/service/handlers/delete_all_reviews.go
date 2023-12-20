package handlers

//import (
//	"github.com/Masterminds/squirrel"
//	"github.com/go-chi/chi"
//	"github.com/jmoiron/sqlx"
//	"net/http"
//	"review_api/internal/data/pg"
//	"review_api/resources"
//	"strconv"
//)
//
//type DeleteReview struct {
//	DB *sqlx.DB
//}
//
//func (h *DeleteReview) DeleteReviews(w http.ResponseWriter, r *http.Request) *resources.APIError {
//	productID, err := strconv.Atoi(chi.URLParam(r, "product_id"))
//	if err != nil {
//		return resources.NewAPIError(http.StatusBadRequest, "StatusBadRequest", "Wrong format product_id")
//	}
//	reviewService := pg.NewReviewService(h.DB)
//	err = reviewService.DeleteReviewsByProductID(productID)
//	if err != nil {
//		return resources.NewAPIError(http.StatusBadGateway, "StatusBadGateway", "Database problem")
//	}
//	// response about successful deleting
//	w.WriteHeader(http.StatusNoContent)
//	return nil
//}
//
//func (s *ReviewService) DeleteReviewsByProductID(productID int) error {
//	builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
//		Delete("").
//		From("reviews").
//		Where(squirrel.Eq{"product_id": productID})
//
//	query, args, err := builder.ToSql()
//	if err != nil {
//		return err
//	}
//	_, err = s.DB.Exec(query, args...)
//	return err
//
//}
