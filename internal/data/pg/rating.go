package pg

import (
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"
	"gitlab.com/distributed_lab/logan/v3/errors"

	"review_api/resources"

	"gitlab.com/distributed_lab/kit/pgdb"

	"review_api/internal/data"
)

const ratingsTableName = "review_ratings"

type ratingQImpl struct {
	db  *pgdb.DB
	sql sq.SelectBuilder
}

func NewRatingQ(db *pgdb.DB) data.RatingQ {
	return &ratingQImpl{
		db:  db.Clone(),
		sql: sq.Select("r.*").From(fmt.Sprintf("%s as r", reviewsTableName)),
	}
}

func (q *ratingQImpl) New() data.RatingQ {
	return NewRatingQ(q.db)
}

func (q *ratingQImpl) Insert(rating data.Rating) error {
	log.Printf("Inserting rating: review_id=%d, user_id=%d, rating=%f", rating.ReviewID, rating.UserID, rating.Rating)

	stmt := sq.Insert(ratingsTableName).
		Columns("review_id", "user_id", "rating").
		Values(rating.ReviewID, rating.UserID, rating.Rating)

	err := q.db.Exec(stmt)
	if err != nil {
		log.Printf("Error inserting rating: %v", err)
		return errors.Wrap(err, "failed to insert rating")
	}
	return nil
}

func (q *ratingQImpl) UpdateRating(ratingID int64, updateData resources.UpdateRatingData) (data.Rating, error) {
	builder := sq.Update(ratingsTableName).Where(sq.Eq{"id": ratingID})

	if updateData.ReviewId != nil {
		builder = builder.Set("review_id", *updateData.ReviewId)
	}
	if updateData.UserId != nil {
		builder = builder.Set("user_id", *updateData.UserId)
	}
	if updateData.Rating != nil {
		builder = builder.Set("rating", *updateData.Rating)
	}

	var updatedRating data.Rating
	err := q.db.Get(&updatedRating, builder)
	return updatedRating, err
}

func (q *ratingQImpl) DeleteRating(ratingID int64) error {
	stmt := sq.Delete(ratingsTableName).Where("id = ?", ratingID)
	return q.db.Exec(stmt)
}

func (q *ratingQImpl) DeleteRatingsByProductID(productID int64) error {
	stmt := sq.Delete(ratingsTableName).
		Where("review_id IN (SELECT id FROM reviews WHERE product_id = ?)", productID)
	err := q.db.Exec(stmt)
	return err
}
