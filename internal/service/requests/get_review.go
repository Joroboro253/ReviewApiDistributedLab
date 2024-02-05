package requests

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/sirupsen/logrus"
	"gitlab.com/distributed_lab/logan/v3/errors"

	"review_api/resources"
)

func NewGetReviewRequest(r *http.Request) (resources.ReviewQueryParams, int64, error) {
	request := resources.ReviewQueryParams{}

	productId, err := strconv.ParseInt(chi.URLParam(r, "product_id"), 10, 64)
	if err != nil {
		logrus.WithError(err).Error("Failed parse Id from request")
		return request, productId, errors.Wrap(err, "failed parse Id from request")
	}

	limitParam := r.URL.Query().Get("limit")
	if limitParam != "" {
		parsedLimit, err := strconv.ParseInt(limitParam, 10, 64)
		if err != nil {
			logrus.WithError(err).Error("Bad limit parameter")
			return request, productId, errors.Wrap(err, "bad limit parameter")
		}
		request.Limit = parsedLimit
	} else {
		request.Limit = 10
	}

	includeRatingsParam := r.URL.Query().Get("includeRatings")
	if includeRatingsParam != "" {
		parsedIncludeRatings, err := strconv.ParseBool(includeRatingsParam)
		if err != nil {
			logrus.WithError(err).Error("Bad include rating parameter")
			return request, productId, errors.Wrap(err, "bad include rating parameter")
		}
		request.IncludeRatings = parsedIncludeRatings
	}

	pageParam := r.URL.Query().Get("page")
	if pageParam != "" {
		parsedPageParam, err := strconv.ParseInt(pageParam, 10, 64)
		if err != nil {
			logrus.WithError(err).Error("Bad page parameter")
			return request, productId, errors.Wrap(err, "bad page parameter")
		}
		request.Page = parsedPageParam
	} else {
		request.Page = 1
	}

	sortByParam := r.URL.Query().Get("sortBy")
	if sortByParam != "" && (sortByParam == "date" || sortByParam == "avgRating" || sortByParam == "productRating") {
		request.SortBy = sortByParam
	} else {
		request.SortBy = "date"
	}

	sortDirectionParam := r.URL.Query().Get("sortDirection")
	if sortDirectionParam != "" && (sortDirectionParam == "asc" || sortDirectionParam == "desc") {
		request.SortDirection = sortDirectionParam
	} else {
		request.SortDirection = "asc"
	}

	if err := ValidateGetReviewParameters(request); err != nil {
		logrus.WithError(err).Error("Validation of get review params failed")
		return request, productId, errors.Wrap(err, "Validation failed")
	}

	return request, productId, nil
}

func ValidateGetReviewParameters(r resources.ReviewQueryParams) error {
	return validation.Errors{
		"includeRatings": validation.Validate(&r.IncludeRatings, validation.In(true, false)),
		"limit":          validation.Validate(&r.Limit, validation.Min(1)),
		"page":           validation.Validate(&r.Limit, validation.Min(1)),
		"sortBy":         validation.Validate(&r.SortBy, validation.In("avgRating", "date", "productRating")),
		"sortDirection":  validation.Validate(&r.SortDirection, validation.In("asc", "desc")),
	}.Filter()
}
