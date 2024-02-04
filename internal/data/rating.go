package data

import (
	"time"

	"review_api/resources"
)

type RatingQ interface {
	New() RatingQ
	Insert(rating Rating) error
	UpdateRating(ratingID int64, updateData resources.UpdateRatingData) error
	DeleteRating(ratingID int64) error
}

type Rating struct {
	ID        int64     `db:"id"`
	ReviewID  int64     `db:"review_id"`
	UserID    int64     `db:"user_id"`
	Rating    int64     `db:"rating"`
	CreatedAt time.Time `db:"created_at"`
	UpdateAt  time.Time `db:"updated_at"`
}
