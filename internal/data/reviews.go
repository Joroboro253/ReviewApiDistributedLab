package data

import (
	"gitlab.com/distributed_lab/logan/v3/errors"
	"time"
)

type ReviewQ interface {
	New() ReviewQ

	Get() (*Review, error)

	Delete(blobId int64) error
	Select() ([]Review, error)

	Transaction(fn func(q ReviewQ) error) error

	Insert(data Review) (Review, error)

	FilterByID(id ...int64) ReviewQ
}

type Review struct {
	ID        int       `json:"id" db:"id"`
	ProductID int       `json:"product_id" db:"product_id"`
	UserID    int       `json:"user_id" db:"user_id"`
	Content   string    `json:"content" db:"content"`
	Rating    float32   `json:"rating" db:"rating"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type ReviewResponse struct {
	Data Review `json:"data"`
}

func (r *Review) Validate() error {
	if r.Rating < 1 || r.Rating > 5 {
		return errors.New("rating must be between 1 and 5")
	}
	return nil
}
