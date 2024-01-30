package requests

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/urlval"

	"review_api/resources"
)

func DeleteReviewRequestByProductID(r *http.Request) (resources.DeleteReviewRequest, error) {
	request := resources.DeleteReviewRequest{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	request.ProductId = cast.ToInt64(chi.URLParam(r, "product_id"))

	return request, nil
}
