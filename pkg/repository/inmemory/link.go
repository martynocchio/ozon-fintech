package inmemory

import (
	"context"
	"fmt"
	ozon_fintech "ozon-fintech"
)

func (r Repository) CreateShortURL(_ context.Context, link *ozon_fintech.Link) (string, error) {
	if _, ok := r.briefToFull[link.Token]; ok {
		return "", fmt.Errorf("token already exist")
	}

	if token, ok := r.fullToBrief[link.BaseURL]; ok {
		return token, nil
	}

	r.briefToFull[link.Token] = link.BaseURL
	r.fullToBrief[link.BaseURL] = link.Token

	return link.Token, nil
}

func (r Repository) GetBaseURL(_ context.Context, link *ozon_fintech.Link) (string, error) {
	if baseURL, ok := r.briefToFull[link.Token]; ok {
		return baseURL, nil
	}

	return "", fmt.Errorf("URL with this token not exist")
}