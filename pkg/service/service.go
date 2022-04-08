package service

import (
	"fmt"
	"math/rand"
	ozon_fintech "ozon-fintech"
	"regexp"
	"time"
)

const pool = "_0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Services interface {
	CreateShortURL(*ozon_fintech.Link) (string, error)
	GetBaseURL(*ozon_fintech.Link) (string, error)
}

func ValidateBaseURL(p *ozon_fintech.Link) error {
	if p == nil {
		return fmt.Errorf("pass nil pointer")
	}

	if p.BaseURL == "" {
		return fmt.Errorf("empty query")
	}

	pattern := `^(https?://|www.)?[a-zA-Z0-9-]{1,256}([.][a-zA-Z-]{1,256})?([.][a-zA-Z]{1,30})([/][a-zA-Z0-9/?=%&#_.-]+)`
	if valid, _ := regexp.Match(pattern, []byte(p.BaseURL)); !valid {
		return fmt.Errorf("%v is a invalid base url", p.BaseURL)
	}

	return nil
}

func ValidateToken(p *ozon_fintech.Link) error {
	if p == nil {
		return fmt.Errorf("pass nil pointer")
	}

	pattern := `^[a-zA-Z0-9_]{10}`
	if valid, _ := regexp.Match(pattern, []byte(p.Token)); !valid {
		return fmt.Errorf("%v is a invalid token", p.Token)
	}

	return nil
}

func GenerateToken() string {
	rand.Seed(time.Now().UnixNano())
	bytes := make([]byte, 10)
	for i := 0; i < 10; i++ {
		bytes[i] = pool[rand.Intn(len(pool))]
	}
	return string(bytes)
}
