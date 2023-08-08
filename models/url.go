package models

type URL struct {
	ID        uint   `json:"id"`
	Original  string `json:"original"`
	ShortCode string `json:"shortCode"`
}
