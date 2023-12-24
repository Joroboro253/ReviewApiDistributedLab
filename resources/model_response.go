package resources

import "review_api/internal/data"

type ResponseData struct {
	Type       string      `json:"type"`
	ID         int         `json:"id"`
	Attributes data.Review `json:"attributes"`
}

type ResponseBody struct {
	Data ResponseData `json:"data"`
}

type SuccessResponse struct {
	Data struct {
		Type       string      `json:"type"`
		ID         int         `json:"id"`
		Attributes interface{} `json:"attributes,omitempty"`
	} `json:"data"`
}

type ReviewListResponse struct {
	Data []data.Review `json:"data"`
}
