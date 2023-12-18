package pg

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"github.com/jmoiron/sqlx"
	"review_api/internal/data"
	"review_api/resources"
)

type ReviewService struct {
	DB *sqlx.DB
}

const reviewsTableName = "reviews"

type reviewRequestsQ struct {
	db  *sqlx.DB
	sql squirrel.SelectBuilder
}

func NewReviewRequestsQ(db *sqlx.DB) data.ReviewRequestsQ {
	return &reviewRequestsQ{
		db:  db,
		sql: squirrel.Select("*").From(fmt.Sprintf("%s as r", reviewsTableName)),
	}
}

func (q *reviewRequestsQ) GetReviewByID(reviewID int) (*resources.Review, error) {
	var result resources.Review
	query, args, err := q.sql.Where(squirrel.Eq{"r.id": reviewID}).ToSql()
	if err != nil {
		return nil, err
	}
	err = q.db.Get(&result, query, args...)
	return &result, err
}

func (q *reviewRequestsQ) InsertReview(value resources.Review) (resources.Review, error) {
	clauses := structs.Map(value)
	var result resources.Review
	stmt, args, err := squirrel.Insert(reviewsTableName).SetMap(clauses).Suffix("returning *").ToSql()
	if err != nil {
		return resources.Review{}, err
	}
	err = q.db.Get(&result, stmt, args...)
	return result, err
}

func (q *reviewRequestsQ) DeleteReviewsByProductID(productID int) error {
	stmt, args, err := squirrel.Delete(reviewsTableName).Where(squirrel.Eq{"product_id": productID}).ToSql()
	if err != nil {
		return err
	}
	_, err = q.db.Exec(stmt, args...)
	return err
}

func NewReviewService(db *sqlx.DB) *ReviewService {
	return &ReviewService{DB: db}
}
