package repository

import (
	"context"
	ozon_fintech "ozon-fintech"
)

type Repository interface {
	CreateShortURL(context.Context, *ozon_fintech.Link) (string, error)
	GetBaseURL(context.Context, *ozon_fintech.Link) (string, error)
}
