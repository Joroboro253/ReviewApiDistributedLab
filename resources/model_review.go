/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "time"

type Review struct {
	// Content of the review
	Content string `json:"Content"`
	// The date and time when the review was created
	CreatedAt time.Time `json:"CreatedAt"`
	// Unique identifier of the review
	Id int64 `json:"Id"`
	// Identifier of the product being reviewed
	ProductId int64 `json:"ProductId"`
	// The date and time when the review was last updated
	UpdatedAt time.Time `json:"UpdatedAt"`
	// Identifier of the user who wrote the review
	UserId int64 `json:"UserId"`
}
