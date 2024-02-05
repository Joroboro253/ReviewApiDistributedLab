package requests

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/spf13/cast"

	"review_api/resources"
)

func DeleteReviewRequest(r *http.Request) (resources.DeleteReviewReq, error) {
	request := resources.DeleteReviewReq{}
	request.ReviewId = cast.ToInt64(chi.URLParam(r, "review_id"))
	return request, nil
}
