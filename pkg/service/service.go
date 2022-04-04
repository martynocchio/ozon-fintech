package service

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

const alphabet = "_0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

type LinkService interface {
	CreateShortURL(ctx context.Context, link *Link) (string, error)
	GetBaseURL(ctx context.Context, link *Link) (string, error)
}

type Link struct {
	BaseURL string `json:"base_url,omitempty" db:"base_url"`
	Token   string `json:"short_url,omitempty" db:"short_url"`
}

func ValidateBaseURL(p *Link) error {
	if p == nil {
		return fmt.Errorf("pass nil pointer")
	}

	if p.BaseURL == "" {
		return fmt.Errorf("empty query")
	}

	pattern := `^(https?://|www.)?[a-zA-Z0-9-]{1,256}([.][a-zA-Z-]{1,256})?([.][a-zA-Z]{1,30})([/]?[a-zA-Z0-9/?=%&#_.-]+)`
	if valid, _ := regexp.Match(pattern, []byte(p.BaseURL)); !valid {
		return fmt.Errorf("%v is a invalid base url", p.BaseURL)
	}

	return nil
}

func ValidateToken(p *Link) error {
	if p == nil {
		return fmt.Errorf("pass nil pointer")
	}

	pattern := `^[a-zA-Z0-9_]{10}$`
	if valid, _ := regexp.Match(pattern, []byte(p.Token)); !valid {
		return fmt.Errorf("%v is a invalid token", p.Token)
	}

	return nil
}

func GenerateToken() string {
	token := bytes.Buffer{}
	uniqueTime := time.Now().Unix()
	_, _ = fmt.Fprintf(&token, "%s", convert(uniqueTime, int64(len(alphabet))))

	for len(token.String()) < 10 {
		rand.Seed(time.Now().UnixNano())
		number := rand.Intn(len(alphabet))
		_, _ = fmt.Fprintf(&token, "%c", alphabet[int64(number)])

	}
	return token.String()
}

func convert(decimalNumber, n int64) string {
	buf := bytes.Buffer{}
	for decimalNumber > 0 {
		curNumber := decimalNumber % n
		decimalNumber /= n
		_, _ = fmt.Fprintf(&buf, "%c", alphabet[curNumber])
	}
	return buf.String()
}

type Error struct {
	Message string `json:"message"`
}

type ValidationError struct {
	Message string      `json:"message"`
	Errors  *ExtraError `json:"errors"`
}

type ExtraError struct {
	AdditionalProperties string `json:"additionalProperties"`
}

func NewError(message string) *Error {
	return &Error{Message: message}
}

func NewValidationError(additionalProperties string) *ValidationError {
	return &ValidationError{
		Message: "invalid data",
		Errors:  NewExtraError(additionalProperties),
	}
}

func NewExtraError(additionalProperties string) *ExtraError {
	return &ExtraError{AdditionalProperties: additionalProperties}
}

type Repository interface {
	CreateShortURL(ctx context.Context, link *Link) (string, error)
	GetBaseURL(ctx context.Context, link *Link) (string, error)
}
type Service struct {
	repos Repository
}

func NewService(repos Repository) *Service {
	return &Service{repos: repos}
}

func (s Service) CreateShortURL(ctx context.Context, link *Link) (string, error) {
	link.Token = GenerateToken()
	token, err := s.repos.CreateShortURL(ctx, link)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s Service) GetBaseURL(ctx context.Context, link *Link) (string, error) {
	baseURL, err := s.repos.GetBaseURL(ctx, link)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}
	return baseURL, nil
}
