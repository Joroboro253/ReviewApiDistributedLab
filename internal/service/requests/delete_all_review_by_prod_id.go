package requests

import (
	"github.com/go-chi/chi"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/urlval"
	"net/http"
)

type DeleteReviewRequestByProdID struct {
	ProductID int64 `url:"-"`
}

func DeleteReviewRequestByProductID(r *http.Request) (DeleteReviewRequestByProdID, error) {
	request := DeleteReviewRequestByProdID{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	request.ProductID = cast.ToInt64(chi.URLParam(r, "product_id"))

	return request, nil
}
