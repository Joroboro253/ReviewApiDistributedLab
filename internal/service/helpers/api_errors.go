package helpers

import (
	"github.com/google/jsonapi"
)

func NewInvalidParamsError() *jsonapi.ErrorObject {
	return &jsonapi.ErrorObject{
		Code:   "INVALID_PARAMS",
		Detail: "One or more of the request parameters are invalid.",
		ID:     "bad_request",
		Status: "400",
		Title:  "Invalid Request Parameters",
	}
}

func NewInternalServerError() *jsonapi.ErrorObject {
	return &jsonapi.ErrorObject{
		Code:   "INTERNAL_SERVER_ERROR",
		Detail: "The server encountered an unexpected condition that prevented it from fulfilling the request",
		ID:     "internal_server",
		Status: "500",
		Title:  "Internal Server Error",
	}
}
