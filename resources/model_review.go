/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import (
	"time"
)

type Review struct {
	Content   string    `json:"Content"`
	CreatedAt time.Time `json:"CreatedAt"`
	ID        int64     `json:"ID"`
	ProductID int64     `json:"ProductID"`
	UpdatedAt time.Time `json:"UpdatedAt"`
	UserID    int64     `json:"UserID"`
}
