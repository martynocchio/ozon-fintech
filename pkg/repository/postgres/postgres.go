package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	ozon_fintech "ozon-fintech"
)

const linksTable = "links"

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (r Repository) CreateShortURL(link *ozon_fintech.Link) (string, error) {
	var token string

	query := fmt.Sprintf("SELECT token FROM %s WHERE base_url = $1", linksTable)
	row := r.db.QueryRow(query, link.BaseURL)

	err := row.Scan(&token)
	if err == nil {
		return token, nil
	}

	query = fmt.Sprintf("INSERT INTO %s (base_url, token) VALUES($1, $2) RETURNING token", linksTable)
	err = r.db.QueryRow(query, link.BaseURL, link.Token).Scan(&token)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r Repository) GetBaseURL(link *ozon_fintech.Link) (string, error) {
	query := fmt.Sprintf("SELECT base_url FROM %s WHERE token = $1", linksTable)
	row := r.db.QueryRow(query, link.Token)

	var baseURL string
	err := row.Scan(&baseURL)

	if err != nil {
		return "", err
	}

	return baseURL, nil
}
