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
			helpers.CtxRatingsQ(pg.NewRatingQ(s.db)),
		),
	)

	r.Route("/products/{product_id}/reviews", func(r chi.Router) {
		r.Post("/", handlers.CreateReview)
		r.Get("/", handlers.GetReviews)
		r.Delete("/", handlers.DeleteReviews)

		r.Route("/{review_id}", func(r chi.Router) {
			r.Patch("/", handlers.UpdateReview)
			r.Post("/", handlers.CreateRating)
			r.Delete("/", handlers.DeleteReview)
			r.Patch("/ratings/{rating_id}", handlers.UpdateRating)
			r.Delete("/ratings/{rating_id}", handlers.DeleteRating)
		})
	})

	return r
}
