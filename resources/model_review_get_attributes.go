/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "time"

type ReviewGetAttributes struct {
	AvgRating float64   `json:"avgRating"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	ReviewId  int64     `json:"reviewId"`
	UpdatedAt time.Time `json:"updatedAt"`
	UserId    int64     `json:"userId"`
}
