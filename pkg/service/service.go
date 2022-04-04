package service

import (
	"bytes"
	"context"
	"fmt"
	"math/rand"
	ozon_fintech "ozon-fintech"
	"regexp"
	"time"
)

const alphabet = "_0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

type LinkService interface {
	CreateShortURL(context.Context, *ozon_fintech.Link) (string, error)
	GetBaseURL(context.Context, *ozon_fintech.Link) (string, error)
}

func ValidateBaseURL(p *ozon_fintech.Link) error {
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

func ValidateToken(p *ozon_fintech.Link) error {
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
