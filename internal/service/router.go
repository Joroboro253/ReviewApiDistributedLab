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
		),
	)

	r.Route("/products/{product_id}/reviews", func(r chi.Router) {
		r.Post("/", handlers.CreateReview)
		r.Get("/", handlers.GetReviews)
		r.Delete("/", handlers.DeleteAllByProductId)
		r.Delete("/{review_id}", handlers.DeleteReviewByID)
		//r.Patch("/{review_id}", helpers.ErrorHandler(reviewHandler.UpdateReviewById))
	})

	return r
}
