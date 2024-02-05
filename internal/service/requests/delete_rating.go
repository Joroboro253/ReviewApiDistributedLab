package requests

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/spf13/cast"

	"review_api/resources"
)

func DeleteRatingRequest(r *http.Request) (resources.DeleteRatingReq, error) {
	request := resources.DeleteRatingReq{}

	request.RatingId = cast.ToInt64(chi.URLParam(r, "rating_id"))

	return request, nil
}
