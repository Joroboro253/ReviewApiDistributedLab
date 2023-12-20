package handlers

//
//import (
//	"github.com/jmoiron/sqlx"
//	"log"
//	"net/http"
//	"review_api/internal/service/helpers"
//	"review_api/resources"
//	"strconv"
//)
//
//type GetReview struct {
//	DB *sqlx.DB
//}
//
//func (h *GetReview) GetReviews(w http.ResponseWriter, r *http.Request) *resources.APIError {
//	handler := helpers.GetHandler{
//		DB: h.DB,
//	}
//	productID, apiErr := helpers.GetProductIDFromURL(r)
//	if apiErr != nil {
//		errorMsg := "Error getting product ID from URL"
//		log.Printf("%s: %v", errorMsg, apiErr)
//		return resources.NewAPIError(http.StatusBadRequest, "StatusBadRequest", errorMsg)
//	}
//	if productID < 1 {
//		errorMsg := "Invalid product_id"
//		log.Printf("%s: %d", errorMsg, productID)
//		return resources.NewAPIError(http.StatusBadRequest, "StatusBadRequest", errorMsg)
//	}
//	sortField, pageStr, limitStr := handler.QueryParameterProcessing(r)
//	// Conversion page and limit
//	page, err := strconv.Atoi(pageStr)
//	if err != nil || page < 1 {
//		page = 1
//	}
//
//	limit, err := strconv.Atoi(limitStr)
//	if err != nil || limit < 1 {
//		limit = 10 // default value
//	}
//	// Pagination
//	_, response, apiErr := handler.Pagination(productID, sortField, strconv.Itoa(page), strconv.Itoa(limit))
//	if apiErr != nil {
//		return apiErr
//	}
//
//	// Generate and send response
//	return handler.GenerateResponse(w, response)
//}

//func (s *ReviewService) GetReviewsByProductID(productID int, sortField string, page, limit int) ([]resources.Review, int, int, error) {
//	// Pagination param check
//	if page < 1 {
//		page = 1
//	}
//	if limit < 1 {
//		limit = 10 // Default setting
//	}
//	// Getting revews
//	countBuilder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
//		Select("COUNT(*)").
//		From("reviews").
//		Where(squirrel.Eq{"product_id": productID})
//
//	// Counting the total number of reviews
//	countQuery, countArgs, err := countBuilder.ToSql()
//	log.Printf("Count Query: %s, Args: %v", countQuery, countArgs)
//	if err != nil {
//		log.Printf("error building count SQL query: %v", err)
//		return nil, 0, 0, fmt.Errorf("error building count SQL query: %w", err)
//	}
//
//	var totalReviews int
//	err = s.DB.Get(&totalReviews, countQuery, countArgs...)
//	if err != nil {
//		log.Printf("error executing count SQL query: %v", err)
//		return nil, 0, 0, fmt.Errorf("error executing count SQL query: %w", err)
//	}
//
//	reviewBuilder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
//		Select("*").
//		From("reviews").
//		Where(squirrel.Eq{"product_id": productID})
//
//	if sortField != "" {
//		reviewBuilder = reviewBuilder.OrderBy(sortField)
//	}
//
//	query, args, err := reviewBuilder.Limit(uint64(limit)).Offset(uint64((page - 1) * limit)).ToSql()
//	if err != nil {
//		log.Printf("error building SQL query: %v", err)
//		return nil, 0, 0, fmt.Errorf("error building SQL query: %w", err)
//	}
//	// Pagination
//	var reviews []resources.Review
//	err = s.DB.Select(&reviews, query, args...)
//	if err != nil {
//		log.Printf("error executing SQL query: %v", err)
//		return nil, 0, 0, fmt.Errorf("error executing SQL query: %w", err)
//	}
//
//	// Calculation of total pages
//	totalPages := int(math.Ceil(float64(totalReviews) / float64(limit)))
//
//	// Return of results
//	return reviews, totalReviews, totalPages, nil
//}
