package requests

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/urlval"

	"review_api/resources"
)

func DeleteRatingRequest(r *http.Request) (resources.DeleteRatingReq, error) {
	request := resources.DeleteRatingReq{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	request.RatingId = cast.ToInt64(chi.URLParam(r, "rating_id"))

	return request, nil
}
