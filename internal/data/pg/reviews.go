package pg

import (
	"fmt"
	"log"

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
		Suffix("RETURNING id, product_id, user_id, content, created_at, updated_at") // ?

	var result data.Review
	err := q.db.Get(&result, stmt)
	if err != nil {
		return errors.Wrap(err, "failed to insert rating")
	}
	return err
}

func (q *reviewQImpl) UpdateReview(reviewID int64, updateData resources.UpdateReviewData) (data.Review, error) {
	builder := sq.Update(reviewsTableName).Where(sq.Eq{"id": reviewID})

	updateFields := false
	if updateData.ProductId != nil {
		builder = builder.Set("product_id", *updateData.ProductId)
		updateFields = true
	}
	if updateData.UserId != nil {
		builder = builder.Set("user_id", *updateData.UserId)
		updateFields = true
	}
	if updateData.Content != nil {
		builder = builder.Set("content", *updateData.Content)
		updateFields = true
	}

	if !updateFields {
		return data.Review{}, errors.New("no fields to update")
	}

	query, args, err := builder.ToSql()

	res, err := q.db.ExecWithResult(sq.Expr(query, args...))

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		return data.Review{}, err
	}

	if rowsAffected == 0 {
		return data.Review{}, errors.New("no rows updated")
	}

	var updatedReview data.Review
	err = q.db.Get(&updatedReview, sq.Select("*").From(reviewsTableName).Where(sq.Eq{"id": reviewID}))
	if err != nil {
		return data.Review{}, err
	}

	return updatedReview, nil
}

func (q *reviewQImpl) Select(sortParam resources.SortParam, includeRatings bool) ([]data.ReviewWithRatings, error) {
	var reviewsWithRatings []data.ReviewWithRatings

	sortFields := map[string]string{
		"date":   "reviews.created_at",
		"rating": "avg_rating",
	}

	selectFields := []string{
		"reviews.id",
		"reviews.product_id",
		"reviews.user_id",
		"reviews.content",
		"reviews.created_at",
		"reviews.updated_at",
	}

	baseQuery := sq.Select(selectFields...).From("reviews")

	if includeRatings {
		baseQuery = baseQuery.
			Column("COALESCE(AVG(review_ratings.rating), 0) AS avg_rating").
			LeftJoin("review_ratings ON reviews.id = review_ratings.review_id").
			GroupBy("reviews.id", "reviews.product_id", "reviews.user_id", "reviews.content", "reviews.created_at", "reviews.updated_at")
	}

	var orderBy string
	if field, ok := sortFields[sortParam.SortBy]; ok {
		if sortParam.SortDirection == "desc" {
			orderBy = fmt.Sprintf("%s DESC", field)
		} else {
			orderBy = fmt.Sprintf("%s ASC", field)
		}
	} else {
		if sortParam.SortDirection == "desc" {
			orderBy = "reviews.created_at DESC"
		} else {
			orderBy = "reviews.created_at ASC"
		}
	}

	query := baseQuery.OrderBy(orderBy).Limit(uint64(sortParam.Limit)).Offset(uint64((sortParam.Page - 1) * sortParam.Limit))

	err := q.db.Select(&reviewsWithRatings, query)
	if err != nil {
		return nil, err
	}

	return reviewsWithRatings, nil
}

func (q *reviewQImpl) DeleteAllByProductId(productId int64) error {
	log.Printf("ID: %d", productId)
	stmt := sq.Delete(reviewsTableName).Where("product_id = ?", productId)
	return q.db.Exec(stmt)
}
