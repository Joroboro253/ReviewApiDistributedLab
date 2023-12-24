package service

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"review_api/internal/data/pg"
	"review_api/internal/service/handlers"
	"review_api/internal/service/helpers"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			helpers.CtxLog(s.log),
			helpers.CtxReviewsQ(pg.NewReviewsQ(s.db)),
			helpers.CtxReviewRequestsQ(pg.NewReviewRequestsQ(s.db)),
		),
	)

	r.Route("/products/{product_id}/reviews", func(r chi.Router) {
		r.Post("/", handlers.CreateReview)
		//r.Get("/", helpers.ErrorHandler(reviewHandler.GetReviews))
		//r.Delete("/", helpers.ErrorHandler(reviewHandler.DeleteReviews))
		//r.Patch("/{review_id}", helpers.ErrorHandler(reviewHandler.UpdateReviewById))
	})

	return r
}
