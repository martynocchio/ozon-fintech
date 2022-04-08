package repository

import (
	ozon_fintech "ozon-fintech"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

type Repository interface {
	CreateShortURL(*ozon_fintech.Link) (string, error)
	GetBaseURL(*ozon_fintech.Link) (string, error)
}
