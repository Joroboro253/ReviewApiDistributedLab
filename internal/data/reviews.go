package data

import (
	"time"

	"review_api/resources"
)

type ReviewQ interface {
	New() ReviewQ
	DeleteAllByProductId(productId int64) error
	DeleteReview(reviewId int64) error
	Select(sortParam resources.SortParam, includeRatings bool, productId int64) ([]ReviewWithRatings, *resources.PaginationMeta, error)
	UpdateReview(updateData resources.UpdateReviewData) error
	Insert(data Review) error
}

type Review struct {
	ID        int64     `db:"id"`
	ProductID int64     `db:"product_id"`
	UserID    int64     `db:"user_id"`
	Content   string    `db:"content"`
	Rating    int64     `db:"rating"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type ReviewWithRatings struct {
	ID           int64     `db:"id"`
	ProductID    int64     `db:"product_id"`
	UserID       int64     `db:"user_id"`
	Content      string    `db:"content"`
	Rating       int64     `db:"rating"`
	RatingsCount int64     `db:"ratings_count"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
	AvgRating    float64   `db:"avg_rating"`
}
