package pg

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"gitlab.com/distributed_lab/logan/v3/errors"

	"gitlab.com/distributed_lab/kit/pgdb"

	"review_api/internal/data"
	"review_api/resources"
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
	stmt := sq.Insert(ratingsTableName).
		Columns("review_id", "user_id", "rating").
		Values(rating.ReviewID, rating.UserID, rating.Rating)

	err := q.db.Exec(stmt)
	if err != nil {
		return errors.Wrap(err, "failed to insert rating")
	}
	return nil
}

func (q *ratingQImpl) UpdateRating(ratingID int64, updateData resources.UpdateRatingData) (data.Rating, error) {
	updateBuilder := sq.Update(ratingsTableName).Where(sq.Eq{"id": ratingID})
	if updateData.Attributes.ReviewId != 0 {
		updateBuilder = updateBuilder.Set("review_id", updateData.Attributes.ReviewId)
	}
	if updateData.Attributes.UserId != 0 {
		updateBuilder = updateBuilder.Set("user_id", updateData.Attributes.UserId)
	}
	if updateData.Attributes.Rating != 0 {
		updateBuilder = updateBuilder.Set("rating", updateData.Attributes.Rating)
	}

	updateSql, args, err := updateBuilder.ToSql()
	if err != nil {
		return data.Rating{}, err
	}

	err = q.db.ExecRaw(updateSql, args...)

	fetchBuilder := sq.Select("*").From(ratingsTableName).Where(sq.Eq{"id": ratingID})
	fetchSql, args, err := fetchBuilder.ToSql()

	var updatedRating data.Rating
	err = q.db.GetRaw(&updatedRating, fetchSql, args...)

	return updatedRating, nil
}
