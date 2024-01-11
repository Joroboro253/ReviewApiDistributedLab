package requests

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/urlval"
)

type DeleteReviewRequestByRevID struct {
	ReviewID int64 `url:"-"`
}

func DeleteReviewRequestByReviewID(r *http.Request) (DeleteReviewRequestByRevID, error) {
	request := DeleteReviewRequestByRevID{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	request.ReviewID = cast.ToInt64(chi.URLParam(r, "review_id"))

	return request, nil
}
