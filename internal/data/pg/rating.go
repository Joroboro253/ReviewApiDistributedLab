package pg

import (
	sq "github.com/Masterminds/squirrel"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"review_api/internal/data"
)

const ratingsTableName = "review_ratings"

type ratingQImpl struct {
	db *pgdb.DB
}

func (q *ratingQImpl) DeleteRatingsByProductID(productID int64) error {
	stmt := sq.Delete(ratingsTableName).
		Where("review_id IN (SELECT id FROM reviews WHERE product_id = ?)", productID)
	err := q.db.Exec(stmt)
	return err
}

func NewRatingQ(db *pgdb.DB) data.RatingQ {
	return &ratingQImpl{db: db}
}

func (q *ratingQImpl) New() data.RatingQ {
	return NewRatingQ(q.db)
}

func (q *ratingQImpl) Insert(rating data.Rating) (data.Rating, error) {
	if rating.Rating < 1.0 || rating.Rating > 5.0 {
		return data.Rating{}, errors.New("rating must be between 1.0 and 5.0")
	}

	var reviewExists bool
	err := q.db.GetRaw(&reviewExists, "SELECT EXISTS(SELECT 1 FROM reviews WHERE id = $1)", rating.ReviewID)
	if err != nil {
		return data.Rating{}, errors.Wrap(err, "failed to check if review exists")
	}
	if !reviewExists {
		return data.Rating{}, errors.New("review does not exist")
	}

	stmt := sq.Insert(ratingsTableName).
		Columns("review_id", "user_id", "rating", "created_at", "updated_at").
		Values(rating.ReviewID, rating.UserID, rating.Rating, sq.Expr("CURRENT_TIMESTAMP"), sq.Expr("CURRENT_TIMESTAMP")).
		Suffix("RETURNING id, review_id, user_id, rating, created_at")

	var newRating data.Rating
	err = q.db.Get(&newRating, stmt)
	if err != nil {
		return data.Rating{}, errors.Wrap(err, "failed to insert rating")
	}
	return newRating, nil
}

func (q *ratingQImpl) UpdateRating(ratingID int64, updateData map[string]interface{}) (data.Rating, error) {
	if len(updateData) == 0 {
		return data.Rating{}, errors.New("No data to update")
	}

	stmt := sq.Update(ratingsTableName).
		SetMap(updateData).
		Where(sq.Eq{"id": ratingID}).
		Suffix("RETURNING id, review_id, user_id, rating, created_at, updated_at")

	var updatedRating data.Rating
	err := q.db.Get(&updatedRating, stmt)
	if err != nil {
		return data.Rating{}, err
	}
	return updatedRating, nil
}

func (q *ratingQImpl) DeleteRating(ratingID int64) error {
	stmt := sq.Delete(ratingsTableName).Where("id = ?", ratingID)
	err := q.db.Exec(stmt)
	return err
}

func (q *ratingQImpl) Transaction(fn func(q data.RatingQ) error) error {
	return q.db.Transaction(func() error {
		return fn(q)
	})
}
