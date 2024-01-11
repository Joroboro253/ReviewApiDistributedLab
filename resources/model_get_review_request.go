/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type GetReviewRequest struct {
	IncludeRatings bool   `json:"include_ratings"`
	Limit          int64  `json:"limit"`
	Page           int64  `json:"page"`
	ReviewId       int64  `json:"review_id"`
	SortBy         string `json:"sort_by"`
}
