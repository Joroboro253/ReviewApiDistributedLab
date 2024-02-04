package requests

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/spf13/cast"

	"review_api/resources"
)

func DeleteReviewRequestByProductID(r *http.Request) (resources.DeleteReviewRequest, error) {
	request := resources.DeleteReviewRequest{}

	request.ProductId = cast.ToInt64(chi.URLParam(r, "product_id"))

	return request, nil
}
