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

var sortFields = map[string]string{
	"date":   "reviews.created_at",
	"rating": "avg_rating",
}

var selectFields = []string{
	"reviews.id",
	"reviews.product_id",
	"reviews.user_id",
	"reviews.content",
	"reviews.created_at",
	"reviews.updated_at",
}

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
		Columns("product_id", "user_id", "content").
		Values(review.ProductID, review.UserID, review.Content)

	err := q.db.Exec(stmt)
	if err != nil {
		return errors.Wrap(err, "failed to insert rating")
	}
	return nil
}

func (q *reviewQImpl) UpdateReview(reviewID int64, updateData resources.UpdateReviewData) (data.Review, error) {
	updateBuilder := sq.Update(reviewsTableName).Where(sq.Eq{"id": reviewID})

	if updateData.Attributes.ProductId != 0 {
		updateBuilder = updateBuilder.Set("product_id", updateData.Attributes.ProductId)
	}
	if updateData.Attributes.UserId != 0 {
		updateBuilder = updateBuilder.Set("user_id", updateData.Attributes.UserId)
	}
	if updateData.Attributes.Content != "" {
		updateBuilder = updateBuilder.Set("content", updateData.Attributes.Content)
	}

	err := q.db.Exec(updateBuilder)
	if err != nil {
		log.Printf("Error executing querry")
		return data.Review{}, err
	}

	var updatedReview data.Review

	return updatedReview, nil
}

func (q *reviewQImpl) Select(sortParam resources.SortParam, includeRatings bool) ([]data.ReviewWithRatings, *resources.PaginationMeta, error) {
	var reviewsWithRatings []data.ReviewWithRatings
	log.Printf("Sorting params in pg: %+v", sortParam)
	// Getting amount of reviews for metadata
	var totalCount int64
	countQuery := sq.Select("COUNT(*)").From("reviews")
	err := q.db.Get(&totalCount, countQuery)
	if err != nil {
		return nil, nil, err
	}

	meta := &resources.PaginationMeta{
		CurrentPage:  sortParam.Page,
		ItemsPerPage: sortParam.Limit,
		TotalItems:   totalCount,
		TotalPages:   (totalCount + sortParam.Limit - 1) / sortParam.Limit,
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
		orderBy = fmt.Sprintf("%s %s", field, sortParam.SortDirection)
	} else {
		orderBy = fmt.Sprintf("reviews.created_at %s", sortParam.SortDirection)
	}

	query := baseQuery.OrderBy(orderBy).Limit(uint64(sortParam.Limit)).Offset(uint64((sortParam.Page - 1) * sortParam.Limit))

	err = q.db.Select(&reviewsWithRatings, query)
	if err != nil {
		return nil, nil, err
	}

	return reviewsWithRatings, meta, nil
}

func (q *reviewQImpl) DeleteAllByProductId(productId int64) error {
	log.Printf("ID: %d", productId)
	stmt := sq.Delete(reviewsTableName).Where("product_id = ?", productId)
	return q.db.Exec(stmt)
}
