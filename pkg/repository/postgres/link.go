package postgres

import (
	"context"
	"github.com/Masterminds/squirrel"
	ozon_fintech "ozon-fintech"
)

func (r Repository) CreateShortURL(ctx context.Context, link *ozon_fintech.Link) (string, error) {
	var (
		query string
		args  []interface{}
		err   error
		token string
	)

	query, args, err = squirrel.Select("token").
		From(linksTable).
		Where(squirrel.Eq{
			"base_url": link.BaseURL,
		}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return "", err
	}

	if err := r.db.GetContext(ctx, &token, query, args...); err == nil {
		return token, nil
	}

	query, args, err = squirrel.Insert(linksTable).
		SetMap(linkData(link)).
		Suffix("RETURNING token").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return "", err
	}

	if err := r.db.GetContext(ctx, &token, query, args...); err != nil {
		return "", err
	}

	return token, nil
}

func (r Repository) GetBaseURL(ctx context.Context, link *ozon_fintech.Link) (string, error) {
	query, args, err := squirrel.Select("base_url").
		From(linksTable).
		Where(squirrel.Eq{
			"token": link.Token,
		}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return "", err
	}

	var baseURL string
	if err := r.db.GetContext(ctx, &baseURL, query, args...); err != nil {
		return "", err
	}

	return baseURL, nil
}

func linkData(p *ozon_fintech.Link) map[string]interface{} {
	data := map[string]interface{}{
		"base_url": p.BaseURL,
		"token":    p.Token,
	}

	return data
}
