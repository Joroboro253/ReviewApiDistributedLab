package service

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"review_api/internal/service/handlers"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
		),
	)
	r.Route("/products/{product_id}/reviews", func(r chi.Router) {
		r.Post("/", helpers.ErrorHandler(reviewHandler.CreateReview))
		r.Get("/", helpers.ErrorHandler(reviewHandler.GetReviews))
		r.Delete("/", helpers.ErrorHandler(reviewHandler.DeleteReviews))
		r.Patch("/{review_id}", helpers.ErrorHandler(reviewHandler.UpdateReviewById))
	})

	return r
}
