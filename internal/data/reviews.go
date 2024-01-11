package data

import (
	"time"

	"review_api/resources"
)

type ReviewQ interface {
	New() ReviewQ

	Get(reviewID int64) (*Review, error)
	DeleteAllByProductId(reviewId int64) error
	DeleteByReviewId(reviewId int64) error
	Select(sortParam resources.SortParam, includeRatings bool) ([]ReviewWithRatings, error)
	UpdateReview(reviewID int64, updateData resources.UpdateReviewData) (Review, error)
	Insert(data Review) error
}

type Review struct {
	ID        int64     `json:"id" db:"id"`
	ProductID int64     `json:"product_id" db:"product_id"`
	UserID    int64     `json:"user_id" db:"user_id"`
	Content   string    `json:"content" db:"content"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type ReviewResponse struct {
	Data Review `json:"data"`
}

type ReviewWithRatings struct {
	ID        int64     `json:"id" db:"id"`
	ProductID int64     `json:"product_id" db:"product_id"`
	UserID    int64     `json:"user_id" db:"user_id"`
	Content   string    `json:"content" db:"content"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Rating    float64   `json:"rating" db:"rating"`
}

//type ReviewWithRatings struct {
//	Type         string `json:"type"`
//	ID           int64  `json:"id"`
//	ProductID    int64  `json:"product_id" db:"product_id"`
//	UserID       int64  `json:"user_id" db:"user_id"` // Добавленное поле
//	Attributes   Review `json:"attributes"`
//	Relationship struct {
//		Ratings []Rating `json:"ratings"`
//	} `json:"relationships"`
//}
