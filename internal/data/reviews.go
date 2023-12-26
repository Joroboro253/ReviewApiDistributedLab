package data

import (
	"time"
)

type ReviewQ interface {
	New() ReviewQ

	Get(reviewID int64) (*Review, error)
	DeleteAllByProductId(reviewId int64) error
	DeleteByReviewId(reviewId int64) error
	Select(sortBy string, page, limit int) ([]Review, error)
	Update(reviewID int64, updateData map[string]interface{}) (Review, error)
	Transaction(fn func(q ReviewQ) error) error

	Insert(data Review) (Review, error)

	FilterByID(id ...int64) ReviewQ
}

type Review struct {
	ID        int       `json:"id" db:"id"`
	ProductID int       `json:"product_id" db:"product_id"`
	UserID    int       `json:"user_id" db:"user_id"`
	Content   string    `json:"content" db:"content"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type ReviewResponse struct {
	Data Review `json:"data"`
}
