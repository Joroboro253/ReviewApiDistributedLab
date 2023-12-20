package resources

type UpdateRequest struct {
	Data ReviewData `json:"data"`
}

type ReviewRequest struct {
	Data struct {
		Type       string `json:"type"`
		Attributes Review `json:"attributes"`
	} `json:"data"`
}
