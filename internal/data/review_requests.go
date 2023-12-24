package data

// Описываем все методы, которые реализованы в review_request
type ReviewRequestsQ interface {
	New() ReviewRequestsQ
	GetAllReviews(reviewID int) (*Review, error)
	CreateReview(value Review) (Review, error)
	//UpdateReview(value resources.Review)
	DeleteReviewsByProductID(productID int) error
	DeleteReviewById(reviewID int) error
}

type ReviewRequest struct {
	Data struct {
		Type       string `json:"type"`
		Attributes Review `json:"attributes"`
	} `json:"data"`
}
