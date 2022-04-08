package inmemory

import (
	"fmt"
	ozon_fintech "ozon-fintech"
	"sync"
)

type Repository struct {
	mu          sync.Mutex
	briefToFull map[string]string
	fullToBrief map[string]string
}

func NewRepository() *Repository {
	briefToFull := make(map[string]string)
	fullToBrief := make(map[string]string)
	return &Repository{
		briefToFull: briefToFull,
		fullToBrief: fullToBrief,
	}
}

func (r *Repository) CreateShortURL(link *ozon_fintech.Link) (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if token, ok := r.fullToBrief[link.BaseURL]; ok {
		return token, nil
	}

	r.briefToFull[link.Token] = link.BaseURL
	r.fullToBrief[link.BaseURL] = link.Token

	return link.Token, nil
}

func (r *Repository) GetBaseURL(link *ozon_fintech.Link) (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if baseURL, ok := r.briefToFull[link.Token]; ok {
		return baseURL, nil
	}
	return "", fmt.Errorf("URL with this token not exist")
}
