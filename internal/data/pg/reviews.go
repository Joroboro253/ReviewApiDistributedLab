package pg

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/sirupsen/logrus"

	"review_api/resources"

	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"

	"review_api/internal/data"
)

var sortFields = map[string]string{
	"date":          "reviews.created_at",
	"avgRating":     "avg_rating",
	"productRating": "reviews.rating",
}

var selectFields = []string{
	"reviews.id",
	"reviews.product_id",
	"reviews.rating",
	"reviews.user_id",
	"reviews.content",
	"reviews.created_at",
	"reviews.updated_at",
}

const reviewsTableName = "reviews"

type reviewQImpl struct {
	db *pgdb.DB
}

func NewReviewsQ(db *pgdb.DB) data.ReviewQ {
	return &reviewQImpl{
		db: db.Clone(),
	}
}

func (q *reviewQImpl) New() data.ReviewQ {
	return NewReviewsQ(q.db)
}

func (q *reviewQImpl) Insert(review data.Review) error {
	stmt := sq.Insert(reviewsTableName).
		Columns("product_id", "user_id", "content", "rating").
		Values(review.ProductID, review.UserID, review.Content, review.Rating)

	err := q.db.Exec(stmt)
	if err != nil {
		logrus.WithError(err).Error("Failed to prepare SQL for inserting review")
		return errors.Wrap(err, "failed to insert review")
	}
	return nil
}

func (q *reviewQImpl) UpdateReview(updateData resources.UpdateReviewData) error {
	updateBuilder := sq.Update(reviewsTableName).Where(sq.Eq{"id": updateData.Id})

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
		logrus.WithError(err).Error("Failed to execute update review query")
		return err
	}
	return err
}

func (q *reviewQImpl) Select(sortParam resources.SortParam, includeRatings bool, productId int64) ([]data.ReviewWithRatings, *resources.PaginationMeta, error) {
	var reviewsWithRatings []data.ReviewWithRatings

	// Getting amount of reviews for metadata
	var totalCount int64
	countQuery := sq.Select("COUNT(*)").From(reviewsTableName).Where(sq.Eq{"product_id": productId})
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

	baseQuery := sq.Select(selectFields...).From("reviews").Where(sq.Eq{"product_id": productId})

	if includeRatings {
		baseQuery = baseQuery.
			Column("COALESCE(AVG(review_ratings.rating), 0) AS avg_rating").
			Column("COUNT(review_ratings.rating) AS ratings_count").
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
		logrus.WithError(err).Error("Failed to select reviews")
		return nil, nil, err
	}
	logrus.Infof("Successfully selected reviews with ratings for product id: %d", productId)
	return reviewsWithRatings, meta, nil
}

func (q *reviewQImpl) DeleteAllByProductId(productId int64) error {
	stmt := sq.Delete(reviewsTableName).Where("product_id = ?", productId)
	return q.db.Exec(stmt)
}

func (q *reviewQImpl) DeleteReview(reviewId int64) error {
	stmt := sq.Delete(reviewsTableName).Where("id = ?", reviewId)
	return q.db.Exec(stmt)
}
