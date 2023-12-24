package helpers

import (
	"context"
	"gitlab.com/distributed_lab/logan/v3"
	"net/http"
	"review_api/internal/data"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	reviewsQCtxKey
	reviewRequestsQCtxKey
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

func CtxReviewRequestsQ(entry data.ReviewRequestsQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, reviewRequestsQCtxKey, entry)
	}
}

func ReviewsQ(r *http.Request) data.ReviewQ {
	return r.Context().Value(reviewsQCtxKey).(data.ReviewQ).New()
}

func BReviewRequestsQ(r *http.Request) data.ReviewRequestsQ {
	return r.Context().Value(reviewRequestsQCtxKey).(data.ReviewRequestsQ).New()
}