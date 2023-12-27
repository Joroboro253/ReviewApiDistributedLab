package helpers

import (
	"context"
	"gitlab.com/distributed_lab/logan/v3"
	"log"
	"net/http"
	"review_api/internal/data"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	reviewsQCtxKey
	ratingsQCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxReviewsQ(entry data.ReviewQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, reviewsQCtxKey, entry)
	}
}

func ReviewsQ(r *http.Request) data.ReviewQ {
	return r.Context().Value(reviewsQCtxKey).(data.ReviewQ).New()
}

func CtxRatingsQ(entry data.RatingQ) func(ctx context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, ratingsQCtxKey, entry)
	}
}

func RatingsQ(r *http.Request) data.RatingQ {
	if val := r.Context().Value(ratingsQCtxKey); val != nil {
		if rq, ok := val.(data.RatingQ); ok {
			return rq.New()
		}
	}
	log.Printf("Error with ratingQ")
	return nil
}
