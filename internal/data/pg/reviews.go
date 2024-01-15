package pg

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"gitlab.com/distributed_lab/logan/v3"

	"review_api/internal/service/helpers"
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
		log.Println("No fields to update")
		return data.Review{}, errors.New("no fields to update")
	}

	query, args, err := builder.ToSql()
	if err != nil {
		log.Printf("Error building SQL query: %v", err)
		return data.Review{}, err
	}

	res, err := q.db.ExecWithResult(sq.Expr(query, args...))
	if err != nil {
		log.Printf("Error executing SQL query: %v", err)
		return data.Review{}, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		return data.Review{}, err
	}

	if rowsAffected == 0 {
		log.Println("No rows updated")
		return data.Review{}, errors.New("no rows updated")
	}

	var updatedReview data.Review
	err = q.db.Get(&updatedReview, sq.Select("*").From(reviewsTableName).Where(sq.Eq{"id": reviewID}))
	if err != nil {
		log.Printf("Error fetching updated review: %v", err)
		return data.Review{}, err
	}

	return updatedReview, nil
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

func (q *reviewQImpl) Select(r *http.Request, sortParam resources.SortParam, includeRatings bool) ([]data.ReviewWithRatings, error) {
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

	helpers.Log(r).WithFields(logan.F{
		"sortParam":      sortParam,
		"includeRatings": includeRatings,
	}).Info("Executing Select query")

	var orderBy string
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

	query := baseQuery.OrderBy(orderBy).Limit(uint64(sortParam.Limit)).Offset(uint64((sortParam.Page - 1) * sortParam.Limit))

	err := q.db.Select(&reviewsWithRatings, query)
	if err != nil {
		return nil, err
	}

	return reviewsWithRatings, nil
}
