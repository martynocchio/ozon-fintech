package ozon_fintech

type Link struct {
	BaseURL string `json:"base_url" db:"base_url"`
	Token   string `json:"short_url" db:"short_url"`
}
