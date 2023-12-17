package resources

import "time"

// structure of entity
type Review struct {
	ID        int       `json:"id" db:"id"`
	ProductID int       `json:"product_id" db:"product_id" validate:"required,gte=1"`
	UserID    int       `json:"user_id" db:"user_id" validate:"required,gte=1"`
	Content   string    `json:"content" db:"content" validate:"required,max=1000"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type ReviewData struct {
	Type       string `json:"type,omitempty"`
	Attributes Review `json:"attributes,omitempty"`
}

type ReviewUpdate struct {
	UserID *int `json:"user_id"`
	//Rating  *float64 `json:"rating,omitempty"`
	Content *string `json:"content,omitempty"`
}

type ReviewUpdateRequest struct {
	Data struct {
		Type       string       `json:"type"`
		Attributes ReviewUpdate `json:"attributes"`
	} `json:"data"`
}
