package pg

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"review_api/internal/data"
)

const reviewsTableName = "reviews"

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

func (q *reviewQImpl) Update(reviewID int64, updateData map[string]interface{}) (data.Review, error) {
	if len(updateData) == 0 {
		return data.Review{}, errors.New("no data to update")
	}

	stmt := sq.Update(reviewsTableName).
		SetMap(updateData).
		Where(sq.Eq{"id": reviewID}).
		Suffix("RETURNING id, product_id, user_id, content, created_at, updated_at")

	var updatedReview data.Review
	err := q.db.Get(&updatedReview, stmt)
	if err != nil {
		return data.Review{}, err
	}
	return updatedReview, nil
}

func (q *reviewQImpl) Get(reviewID int64) (*data.Review, error) {
	var review data.Review

	query, args, err := sq.Select("*").
		From("reviews").
		Where(sq.Eq{"id": reviewID}).
		ToSql()
	if err != nil {
		return nil, err
	}

	err = q.db.GetRaw(&review, query, args...)
	if err != nil {
		return nil, err
	}

	return &review, nil
}

func (q *reviewQImpl) DeleteByReviewId(reviewId int64) error {
	stmt := sq.Delete(reviewsTableName).Where("id = ?", reviewId)
	err := q.db.Exec(stmt)
	return err
}

func (q *reviewQImpl) DeleteAllByProductId(reviewId int64) error {
	stmt := sq.Delete(reviewsTableName).Where("product_id = ?", reviewId)
	err := q.db.Exec(stmt)
	return err
}

func (q *reviewQImpl) Select(sortBy string, page, limit int, includeRatings bool) ([]data.ReviewWithRatings, error) {
	var reviewsWithRatings []data.ReviewWithRatings

	offset := (page - 1) * limit
	query := q.sql.Limit(uint64(limit)).Offset(uint64(offset))
	switch sortBy {
	case "date":
		query = query.OrderBy("created_at DESC")
	case "rating":
		query = query.OrderBy("rating DESC")
	case "has_rating":
		query = query.OrderBy("CASE WHEN rating IS NOT NULL THEN 1 ELSE 0 END DESC, rating DESC")
	default:
		query = query.OrderBy("created_at DESC")
	}
	// Sort
	// Pagination
	if page > 0 && limit > 0 {
		offset := (page - 1) * limit
		query = query.Limit(uint64(limit)).Offset(uint64(offset))
	}

	var reviews []data.Review
	err := q.db.Select(&reviews, query)
	if err != nil {
		return nil, err
	}

	for _, review := range reviews {
		reviewWithRatings := data.ReviewWithRatings{Review: review}
		if includeRatings {
			ratingsQuery := sq.Select("*").From("review_ratings").Where(sq.Eq{"review_id": review.ID})
			var ratings []data.Rating
			err := q.db.Select(&ratings, ratingsQuery)
			if err != nil {
				return nil, err
			}
			reviewWithRatings.Ratings = ratings
		}
		reviewsWithRatings = append(reviewsWithRatings, reviewWithRatings)
	}

	return reviewsWithRatings, nil
}

func (q *reviewQImpl) Transaction(fn func(q data.ReviewQ) error) error {
	return q.db.Transaction(func() error {
		return fn(q)
	})
}

func (q *reviewQImpl) Insert(review data.Review) (data.Review, error) {
	// Explicitly specifying the columns and values to ensure correct mapping
	stmt := sq.Insert(reviewsTableName).
		Columns("product_id", "user_id", "content", "created_at", "updated_at").
		Values(review.ProductID, review.UserID, review.Content, sq.Expr("CURRENT_TIMESTAMP"), sq.Expr("CURRENT_TIMESTAMP")).
		Suffix("RETURNING id, product_id, user_id, content, created_at, updated_at")

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
