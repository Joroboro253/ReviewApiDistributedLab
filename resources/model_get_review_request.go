/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type GetReviewRequest struct {
	IncludeRatings bool   `json:"includeRatings"`
	Limit          int64  `json:"limit"`
	Page           int64  `json:"page"`
	ReviewId       int64  `json:"reviewId"`
	SortBy         string `json:"sortBy"`
	SortDirection  string `json:"sortDirection"`
}
