package data

import "review_api/resources"

// Описываем все методы, которые реализованы в review_request
type ReviewRequestsQ interface {
	CreateReview(review *resources.Review) (int, *resources.APIError)
	DeleteReviewsByProductID(productID int) error
	GetReviewsByProductID(productID int, sortField string, page, limit int) ([]resources.Review, int, int, error)
	UpdateReview(productId, reviewId int, updateData resources.ReviewUpdate) (int, error)
}
