package helpers

import (
	"review_api/internal/data"
	"review_api/resources"
)

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
