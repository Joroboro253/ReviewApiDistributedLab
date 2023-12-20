package helpers

import (
	"review_api/resources"

	//"github.com/Joroboro253/ReviewApiDistributedLab/internal/service/handlers"
	"net/http"
)

type ReviewHandlerInterface interface {
	GetProductIDFromURL(r *http.Request) (int, *resources.APIError)
	DecodeRequestBody(r *http.Request) (resources.UpdateRequest, *resources.APIError)
	GenerateResponse(w http.ResponseWriter, review resources.Review) *resources.APIError
	CreateReview(productId int, reqBody resources.UpdateRequest) (resources.Review, *resources.APIError)
}
