/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import (
	"time"
)

type Rating struct {
	CreatedAt time.Time `json:"createdAt"`
	Id        int64     `json:"id"`
	Rating    float64   `json:"rating"`
	ReviewID  int64     `json:"reviewID"`
	UpdatedAt time.Time `json:"updatedAt"`
	UserID    int64     `json:"userID"`
}
