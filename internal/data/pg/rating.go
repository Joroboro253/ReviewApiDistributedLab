package pg

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/sirupsen/logrus"
	"gitlab.com/distributed_lab/logan/v3/errors"

	"gitlab.com/distributed_lab/kit/pgdb"

	"review_api/internal/data"
	"review_api/resources"
)

const ratingsTableName = "review_ratings"

type ratingQImpl struct {
	db *pgdb.DB
}

func NewRatingQ(db *pgdb.DB) data.RatingQ {
	return &ratingQImpl{
		db: db.Clone(),
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
		logrus.WithError(err).Error("Failed to insert rating into database")
		return errors.Wrap(err, "failed to insert rating")
	}
	return nil
}

func (q *ratingQImpl) UpdateRating(ratingID int64, updateData resources.UpdateRatingData) error {
	updateBuilder := sq.Update(ratingsTableName).Where(sq.Eq{"id": ratingID})

	if updateData.Attributes.UserId != 0 {
		updateBuilder = updateBuilder.Set("user_id", updateData.Attributes.UserId)
	}
	if updateData.Attributes.Rating != 0 {
		updateBuilder = updateBuilder.Set("rating", updateData.Attributes.Rating)
	}

	err := q.db.Exec(updateBuilder)
	if err != nil {
		logrus.WithError(err).Error("Failed to execute query")
		return err
	}

	return nil
}

func (q *ratingQImpl) DeleteRating(ratingID int64) error {
	stmt := sq.Delete(ratingsTableName).Where("id = ?", ratingID)
	err := q.db.Exec(stmt)
	if err != nil {
		logrus.WithError(err).Error("Failed to execute delete rating query")
		return err
	}

	return nil
}
