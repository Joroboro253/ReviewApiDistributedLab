package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

	"review_api/internal/data"
	"review_api/internal/service/helpers"
	"review_api/internal/service/requests"
	"review_api/resources"
)

func GetReviews(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetReviewRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	reviewQ := helpers.ReviewsQ(r)
	sortParam := resources.SortParam{Limit: request.Limit, Page: request.Page, SortBy: request.SortBy, SortDirection: request.SortDirection}
	helpers.Log(r).WithError(err).Debugf("Sorting params in handler: %+v", sortParam)
	reviews, meta, err := reviewQ.Select(sortParam, request.IncludeRatings)
	if err != nil {
		helpers.Log(r).WithError(err).Info("Internal server Error")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	apiResponse := ConvertToAPIResponse(reviews, meta)
	ape.Render(w, apiResponse)
}

func ConvertToAPIResponse(reviews []data.ReviewWithRatings, meta *resources.PaginationMeta) resources.ReviewApiResponse {
	var apiResponse resources.ReviewApiResponse
	for _, review := range reviews {
		resource := resources.ReviewResource{
			Type:      "reviews",
			ProductId: review.ProductID,
			Attributes: resources.ReviewGetAttributes{
				ReviewId:  review.ID,
				UserId:    review.UserID,
				Content:   review.Content,
				CreatedAt: review.CreatedAt,
				UpdatedAt: review.UpdatedAt,
				AvgRating: review.AvgRating,
			},
		}
		apiResponse.Data = append(apiResponse.Data, resource)
	}
	apiResponse.Meta = meta
	return apiResponse
}
