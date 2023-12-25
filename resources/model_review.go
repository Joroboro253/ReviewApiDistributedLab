package resources

import (
	"review_api/internal/data"
)

// structure of entity

type ReviewData struct {
	Type       string      `json:"type,omitempty"`
	Attributes data.Review `json:"attributes,omitempty"`
}

type ReviewUpdate struct {
	UserID  *int     `json:"user_id"`
	Rating  *float64 `json:"rating,omitempty"`
	Content *string  `json:"content,omitempty"`
}

type ReviewUpdateRequest struct {
	Data struct {
		Type       string       `json:"type"`
		Attributes ReviewUpdate `json:"attributes"`
	} `json:"data"`
}

type UpdateReviewRequest struct {
	Data struct {
		ProductID *int     `json:"product_id,omitempty"`
		UserID    *int     `json:"user_id,omitempty"`
		Content   *string  `json:"content,omitempty"`
		Rating    *float32 `json:"rating,omitempty"`
	} `json:"data"`
}
