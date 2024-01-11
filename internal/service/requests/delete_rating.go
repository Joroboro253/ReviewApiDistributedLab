package requests

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/urlval"
)

type DeleteRatingReq struct {
	RatingID int64 `url:"-"`
}

func DeleteRatingRequest(r *http.Request) (DeleteRatingReq, error) {
	request := DeleteRatingReq{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	request.RatingID = cast.ToInt64(chi.URLParam(r, "rating_id"))

	return request, nil
}
