package pg

import (
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"

	"review_api/resources"

	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"

	"review_api/internal/data"
)

const reviewsTableName = "reviews"

type reviewQImpl struct {
	db  *pgdb.DB
	sql sq.SelectBuilder
}

func NewReviewsQ(db *pgdb.DB) data.ReviewQ {
	return &reviewQImpl{
		db:  db.Clone(),
		sql: sq.Select("r.*").From(fmt.Sprintf("%s as r", reviewsTableName)),
	}
}

func (q *reviewQImpl) New() data.ReviewQ {
	return NewReviewsQ(q.db)
}

func (q *reviewQImpl) Insert(review data.Review) error {
	stmt := sq.Insert(reviewsTableName).
		Columns("product_id", "user_id", "content", "created_at", "updated_at").
		Values(review.ProductID, review.UserID, review.Content, sq.Expr("CURRENT_TIMESTAMP"), sq.Expr("CURRENT_TIMESTAMP")).
		Suffix("RETURNING id, product_id, user_id, content, created_at, updated_at")

	var result data.Review
	err := q.db.Get(&result, stmt)
	if err != nil {
		return errors.Wrap(err, "failed to insert rating")
	}
	return err
}

func (q *reviewQImpl) UpdateReview(reviewID int64, updateData resources.UpdateReviewData) (data.Review, error) {
	builder := sq.Update(reviewsTableName).Where(sq.Eq{"id": reviewID})

	if updateData.ProductId != nil {
		builder = builder.Set("product_id", *updateData.ProductId)
	}
	if updateData.UserId != nil {
		builder = builder.Set("user_id", *updateData.UserId)
	}
	if updateData.Content != nil {
		builder = builder.Set("content", *updateData.Content)
	}

	var updatedReview data.Review
	err := q.db.Get(&updatedReview, builder)
	return updatedReview, err
}

func (q *reviewQImpl) Get(reviewID int64) (*data.Review, error) {
	var review data.Review

	query := sq.Select("*").
		From("reviews").
		Where(sq.Eq{"id": reviewID})

	return &review, q.db.Get(&review, query)
}

func (q *reviewQImpl) DeleteByReviewId(reviewId int64) error {
	stmt := sq.Delete(reviewsTableName).Where("id = ?", reviewId)
	return q.db.Exec(stmt)
}

func (q *reviewQImpl) DeleteAllByProductId(reviewId int64) error {
	stmt := sq.Delete(reviewsTableName).Where("product_id = ?", reviewId)
	return q.db.Exec(stmt)
}

func (q *reviewQImpl) Select(sortParam resources.SortParam, includeRatings bool) ([]data.ReviewWithRatings, error) {
	var reviewsWithRatings []data.ReviewWithRatings

	sortFields := map[string]string{
		"date":   "reviews.created_at",
		"rating": "ratings.rating",
	}

	var orderBy string
	offset := (sortParam.Page - 1) * sortParam.Limit
	baseQuery := sq.Select("reviews.id", "reviews.product_id", "reviews.user_id", "reviews.content", "reviews.created_at", "reviews.updated_at").From("reviews")

	if includeRatings {
		baseQuery = baseQuery.Column("ratings.rating").LeftJoin("ratings ON reviews.id = ratings.review_id")
	}

	if strings.HasPrefix(sortParam.SortBy, "-") {
		sortByField := strings.TrimPrefix(sortParam.SortBy, "-")
		if field, ok := sortFields[sortByField]; ok {
			orderBy = fmt.Sprintf("%s DESC", field)
		} else {
			orderBy = "reviews.created_at DESC"
		}
	} else {
		if field, ok := sortFields[sortParam.SortBy]; ok {
			orderBy = fmt.Sprintf("%s ASC", field)
		} else {
			orderBy = "reviews.created_at ASC"
		}
	}
	query := baseQuery.OrderBy(orderBy).Limit(uint64(sortParam.Limit)).Offset(uint64(offset))

	err := q.db.Select(&reviewsWithRatings, query)
	if err != nil {
		return nil, err
	}

	return reviewsWithRatings, nil
}
