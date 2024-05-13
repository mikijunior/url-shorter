package utils

import (
	"errors"
	"math/rand"
	"net/url"
	"strings"
	"time"
)

const (
	letterBytes    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	shortURLLength = 6
)

var (
	src = rand.NewSource(time.Now().UnixNano())
	r   = rand.New(src)
)

func GenerateShortURL() string {
	b := make([]byte, shortURLLength)
	for i := range b {
		b[i] = letterBytes[r.Intn(len(letterBytes))]
	}
	return string(b)
}

func SanitizeURL(input string) (string, error) {
	if !strings.HasPrefix(input, "http://") && !strings.HasPrefix(input, "https://") {
		input = "https://" + input
	}

	u, err := url.Parse(input)
	if err != nil {
		return "", err
	}

	if u.Host == "" {
		return "", errors.New("invalid URL")
	}

	return u.String(), nil
}
