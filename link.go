package ozon_fintech

type Link struct {
	BaseURL string `json:"base_url,omitempty" db:"base_url"`
	Token   string `json:"short_url,omitempty" db:"short_url"`
}
