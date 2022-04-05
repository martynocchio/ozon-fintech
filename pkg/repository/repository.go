package repository

import (
	"context"
	ozon_fintech "ozon-fintech"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

type Repository interface {
	CreateShortURL(context.Context, *ozon_fintech.Link) (string, error)
	GetBaseURL(context.Context, *ozon_fintech.Link) (string, error)
}
