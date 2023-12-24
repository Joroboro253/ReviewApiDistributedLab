package pg

import (
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
	"review_api/internal/data"
)

const reviewsTableName = "reviews"

func NewReviewRequestsQ(db *pgdb.DB) data.ReviewRequestsQ {
	return &reviewRequestsQ{
		db:  db.Clone(),
		sql: sq.Select("n.*").From(fmt.Sprintf("%s as n", reviewsTableName)),
	}
}

type reviewRequestsQ struct {
	db  *pgdb.DB
	sql sq.SelectBuilder
}

func (q *reviewRequestsQ) New() data.ReviewRequestsQ {
	return NewReviewRequestsQ(q.db)
}

func (q *reviewRequestsQ) GetAllReviews(reviewID int) (*data.Review, error) {
	var result data.Review
	q.sql = q.sql.Where(sq.Eq{"n.id": reviewID})
	err := q.db.Get(&result, q.sql)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &result, err
}

func (q *reviewRequestsQ) CreateReview(value data.Review) (data.Review, error) {
	clauses := structs.Map(value)
	var result data.Review
	stmt := sq.Insert(reviewsTableName).SetMap(clauses).Suffix("returning *")
	err := q.db.Get(&result, stmt)
	return result, err
}

func (q *reviewRequestsQ) DeleteReviewsByProductID(productID int) error {
	stmt := sq.Delete(reviewsTableName).Where(sq.Eq{"product_id": productID})
	err := q.db.Exec(stmt)
	return err
}

func (q *reviewRequestsQ) DeleteReviewById(reviewID int) error {
	stmt := sq.Delete(reviewsTableName).Where(sq.Eq{"id": reviewID})
	err := q.db.Exec(stmt)
	return err
}
