package pg

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"github.com/jmoiron/sqlx"
	"log"
	"math"
	"net/http"
	"review_api/internal/data"
	"review_api/resources"
	"time"
)

type ReviewService struct {
	DB *sqlx.DB
}

const reviewsTableName = "reviews"

type reviewRequestsQ struct {
	db  *sqlx.DB
	sql squirrel.SelectBuilder
}

func NewReviewRequestsQ(db *sqlx.DB) data.ReviewRequestsQInterface {
	return &reviewRequestsQ{
		db:  db,
		sql: squirrel.Select("*").From(fmt.Sprintf("%s as r", reviewsTableName)),
	}
}

func (q *reviewRequestsQ) GetReviewByID(reviewID int) (*resources.Review, error) {
	var result resources.Review
	query, args, err := q.sql.Where(squirrel.Eq{"r.id": reviewID}).ToSql()
	if err != nil {
		return nil, err
	}
	err = q.db.Get(&result, query, args...)
	return &result, err
}

func (q *reviewRequestsQ) InsertReview(value resources.Review) (resources.Review, error) {
	clauses := structs.Map(value)
	var result resources.Review
	stmt, args, err := squirrel.Insert(reviewsTableName).SetMap(clauses).Suffix("returning *").ToSql()
	if err != nil {
		return resources.Review{}, err
	}
	err = q.db.Get(&result, stmt, args...)
	return result, err
}

func (q *reviewRequestsQ) DeleteReviewsByProductID(productID int) error {
	stmt, args, err := squirrel.Delete(reviewsTableName).Where(squirrel.Eq{"product_id": productID}).ToSql()
	if err != nil {
		return err
	}
	_, err = q.db.Exec(stmt, args...)
	return err
}

func NewReviewService(db *sqlx.DB) *ReviewService {
	return &ReviewService{DB: db}
}

func (s *ReviewService) CreateReview(productId int, reviewRequest resources.ReviewRequest) (resources.Review, *resources.APIError) {
	review := resources.Review{
		ProductID: productId,
		UserID:    reviewRequest.Data.Attributes.UserID,
		Content:   reviewRequest.Data.Attributes.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	reviewQ := NewReviewRequestsQ(s.DB)
	createdReview, err := reviewQ.InsertReview(review)
	if err != nil {
		// Преобразуем стандартную ошибку Go в ошибку API
		apiErr := convertToAPIError(err, "Error inserting review into database")
		return resources.Review{}, apiErr
	}

	return createdReview, nil
}

func (s *ReviewService) DeleteReviewsByProductID(productID int) error {
	builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Delete("").
		From("reviews").
		Where(squirrel.Eq{"product_id": productID})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}
	_, err = s.DB.Exec(query, args...)
	return err

}

func (s *ReviewService) GetReviewsByProductID(productID int, sortField string, page, limit int) ([]resources.Review, int, int, error) {
	// Pagination param check
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10 // Default setting
	}
	// Getting revews
	countBuilder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select("COUNT(*)").
		From("reviews").
		Where(squirrel.Eq{"product_id": productID})

	// Counting the total number of reviews
	countQuery, countArgs, err := countBuilder.ToSql()
	log.Printf("Count Query: %s, Args: %v", countQuery, countArgs)
	if err != nil {
		log.Printf("error building count SQL query: %v", err)
		return nil, 0, 0, fmt.Errorf("error building count SQL query: %w", err)
	}

	var totalReviews int
	err = s.DB.Get(&totalReviews, countQuery, countArgs...)
	if err != nil {
		log.Printf("error executing count SQL query: %v", err)
		return nil, 0, 0, fmt.Errorf("error executing count SQL query: %w", err)
	}

	reviewBuilder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select("*").
		From("reviews").
		Where(squirrel.Eq{"product_id": productID})

	if sortField != "" {
		reviewBuilder = reviewBuilder.OrderBy(sortField)
	}

	query, args, err := reviewBuilder.Limit(uint64(limit)).Offset(uint64((page - 1) * limit)).ToSql()
	if err != nil {
		log.Printf("error building SQL query: %v", err)
		return nil, 0, 0, fmt.Errorf("error building SQL query: %w", err)
	}
	// Pagination
	var reviews []resources.Review
	err = s.DB.Select(&reviews, query, args...)
	if err != nil {
		log.Printf("error executing SQL query: %v", err)
		return nil, 0, 0, fmt.Errorf("error executing SQL query: %w", err)
	}

	// Calculation of total pages
	totalPages := int(math.Ceil(float64(totalReviews) / float64(limit)))

	// Return of results
	return reviews, totalReviews, totalPages, nil
}

func (s *ReviewService) UpdateReview(productId, reviewId int, updateData resources.ReviewUpdate) (int, error) {
	// Initialization of SQL-builder queries
	builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Update("reviews")

	// Add Update Conditions if provided
	if updateData.UserID != nil {
		builder = builder.Set("user_id", *updateData.UserID)
	}
	if updateData.Content != nil {
		builder = builder.Set("content", *updateData.Content)
	}

	query, args, err := builder.Where(squirrel.Eq{"id": reviewId, "product_id": productId}).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return 0, err
	}

	// Executing the query
	var updatedReviewID int
	err = s.DB.QueryRow(query, args...).Scan(&updatedReviewID)
	if err != nil {
		return 0, err
	}

	return updatedReviewID, nil
}

func convertToAPIError(err error, message string) *resources.APIError {
	// Здесь может быть логика для определения статуса ошибки и пр.
	return &resources.APIError{
		Status: http.StatusInternalServerError, // Пример статуса
		Title:  "Internal Server Error",
		Detail: message + ": " + err.Error(),
	}
}
