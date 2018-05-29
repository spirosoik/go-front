package front

//Links DTO which includes self url
type Links struct {
	Current string `json:"self"`
}

//Pagination DTO which includes previous and next url
type Pagination struct {
	PrevURL string `json:"prev"`
	NextURL string `json:"next"`
}

//Settings DTO
type Settings struct {
	WebhookURL string `json:"webhook_url,omitempty"`
}
