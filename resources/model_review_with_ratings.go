/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import (
	"time"
)

type ReviewWithRatings struct {
	AvgRating float64   `json:"avg_rating"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	Id        int64     `json:"id"`
	ProductId int64     `json:"product_id"`
	UpdatedAt time.Time `json:"updated_at"`
	UserId    int64     `json:"user_id"`
}
