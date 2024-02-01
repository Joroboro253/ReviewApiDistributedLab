/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import (
	"time"
)

type ReviewDetail struct {
	AvgRating float64   `json:"avgRating"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	UserId    int64     `json:"userId"`
}
