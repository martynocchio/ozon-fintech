package service

import (
	"context"
	"database/sql"
	ozon_fintech "ozon-fintech"
	"ozon-fintech/pkg/repository"
)

type Service struct {
	repos repository.Repository
}

func NewService(repos repository.Repository) *Service {
	return &Service{repos: repos}
}

func (s Service) CreateShortURL(ctx context.Context, link *ozon_fintech.Link) (string, error) {
	link.Token = GenerateToken()
	token, err := s.repos.CreateShortURL(ctx, link)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s Service) GetBaseURL(ctx context.Context, link *ozon_fintech.Link) (string, error) {
	baseURL, err := s.repos.GetBaseURL(ctx, link)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}
	return baseURL, nil
}
