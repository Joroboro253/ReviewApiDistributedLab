package data

import (
	"time"
)

type ReviewQ interface {
	New() ReviewQ

	Get(reviewID int64) (*Review, error)
	DeleteAllByProductId(reviewId int64) error
	DeleteByReviewId(reviewId int64) error
	Select(sortBy string, page, limit int, includeRatings bool) ([]ReviewWithRatings, error)
	Update(reviewID int64, updateData map[string]interface{}) (Review, error)
	Transaction(fn func(q ReviewQ) error) error
	Insert(data Review) (Review, error)

	FilterByID(id ...int64) ReviewQ
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
	Type         string `json:"type"`
	ID           int64  `json:"id"`
	Attributes   Review `json:"attributes"`
	Relationship struct {
		Ratings []Rating `json:"ratings"`
	} `json:"relationships"`
}
