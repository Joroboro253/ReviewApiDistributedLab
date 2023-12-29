package data

import "time"

type RatingQ interface {
	New() RatingQ
	Insert(rating Rating) (Rating, error)
	UpdateRating(ratingID int64, updateData map[string]interface{}) (Rating, error)
	DeleteRating(ratingID int64) error
	Transaction(fn func(q RatingQ) error) error
	DeleteRatingsByProductID(productID int64) error
}

type Rating struct {
	ID        int64     `db:"id"`
	ReviewID  int64     `db:"review_id"`
	UserID    int64     `db:"user_id"`
	Rating    float64   `db:"rating"`
	CreatedAt time.Time `db:"created_at"`
	UpdateAt  time.Time `db:"updated_at"`
}

type RatingResponse struct {
	Data Rating `json:"data"`
}
