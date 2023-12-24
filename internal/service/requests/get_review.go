package requests

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/spf13/cast"

	"gitlab.com/distributed_lab/urlval"
)

type GetReviewRequest struct {
	ReviewID int64 `url:"-"`
}

func NewGetReviewRequest(r *http.Request) (GetReviewRequest, error) {
	request := GetReviewRequest{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	request.ReviewID = cast.ToInt64(chi.URLParam(r, "id"))

	return request, nil
}
