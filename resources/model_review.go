/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import (
	"time"
)

type Review struct {
	// Content of the review
	Content string `json:"Content"`
	// The date and time when the review was created
	CreatedAt *time.Time `json:"CreatedAt,omitempty"`
	// Unique identifier of the review
	ID *int64 `json:"ID,omitempty"`
	// Identifier of the product being reviewed
	ProductId int64 `json:"ProductId"`
	// The date and time when the review was last updated
	UpdatedAt *time.Time `json:"UpdatedAt,omitempty"`
	// Identifier of the user who wrote the review
	UserId int64 `json:"UserId"`
}
