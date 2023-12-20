package helpers

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"review_api/resources"
	"strconv"
)

func GetProductIDFromURL(r *http.Request) (int, *resources.APIError) {
	productIDStr := chi.URLParam(r, "product_id")
	productId, err := strconv.Atoi(productIDStr)
	if err != nil {
		errorMsg := "Wrong format product_id"
		log.Printf("%s: %v", errorMsg, err)
		return 0, resources.NewAPIError(http.StatusBadRequest, "StatusBadRequest", errorMsg)
	}
	return productId, nil
}
