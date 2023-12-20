package data

import (
	"review_api/resources"
)

// Описываем все методы, которые реализованы в review_request
type ReviewRequestsQInterface interface {
	GetReviewByID(reviewID int) (*resources.Review, error)
	InsertReview(value resources.Review) (resources.Review, error)
	DeleteReviewsByProductID(productID int) error
}
