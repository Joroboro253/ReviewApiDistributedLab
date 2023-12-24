package pg

import (
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"gitlab.com/distributed_lab/kit/pgdb"
	"review_api/internal/data"
)

func NewReviewsQ(db *pgdb.DB) data.ReviewQ {
	return &reviewQImpl{
		db:  db.Clone(),
		sql: sq.Select("n.*").From(fmt.Sprintf("%s as n", reviewsTableName)),
	}
}

type reviewQImpl struct {
	db  *pgdb.DB
	sql sq.SelectBuilder
}

func (q *reviewQImpl) New() data.ReviewQ {
	return NewReviewsQ(q.db)
}

func (q *reviewQImpl) Get(reviewId int64) (*data.Review, error) {
	var result data.Review
	q.sql = q.sql.Where(sq.Eq{"n.id": reviewId})
	err := q.db.Get(&result, q.sql)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (q *reviewQImpl) Delete(reviewId int64) error {
	stmt := sq.Delete(reviewsTableName).Where("id = ?", reviewId)
	err := q.db.Exec(stmt)
	return err
}

func (q *reviewQImpl) Select() ([]data.Review, error) {
	var result []data.Review
	err := q.db.Select(&result, q.sql)
	return result, err
}

func (q *reviewQImpl) Transaction(fn func(q data.ReviewQ) error) error {
	return q.db.Transaction(func() error {
		return fn(q)
	})
}

func (q *reviewQImpl) Insert(review data.Review) (data.Review, error) {
	// Explicitly specifying the columns and values to ensure correct mapping
	if err := review.Validate(); err != nil {
		return data.Review{}, err
	}
	stmt := sq.Insert(reviewsTableName).
		Columns("product_id", "user_id", "content", "rating", "created_at", "updated_at").
		Values(review.ProductID, review.UserID, review.Content, review.Rating, sq.Expr("CURRENT_TIMESTAMP"), sq.Expr("CURRENT_TIMESTAMP")).
		Suffix("RETURNING id, product_id, user_id, content, rating, created_at, updated_at")

	var result data.Review
	err := q.db.Get(&result, stmt)
	if err != nil {
		return data.Review{}, err // Return an empty Review and the error
	}

	return result, nil
}

func (q *reviewQImpl) Page(pageParams pgdb.OffsetPageParams) data.ReviewQ {
	q.sql = pageParams.ApplyTo(q.sql, "id")
	return q
}

func (q *reviewQImpl) FilterByID(ids ...int64) data.ReviewQ {
	q.sql = q.sql.Where(sq.Eq{"n.id": ids})
	return q
}
