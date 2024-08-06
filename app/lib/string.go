package lib

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
	"time"
)

func GenerateSlug(name string) string {
	return strings.ToLower(strings.ReplaceAll(name, " ", "-"))
}

func GenerateUniqueSlug(slug string) string {
	return fmt.Sprintf("%s-%d", slug, time.Now().UnixNano())
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateRandomString(n int) (string, error) {
	result := make([]byte, n)
	for i := range result {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		result[i] = letters[num.Int64()]
	}
	return string(result), nil
}

func GenerateOrderNumber() (string, error) {
	part1, err := GenerateRandomString(4)
	if err != nil {
		return "", err
	}
	part2, err := GenerateRandomString(4)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s-%s", part1, part2), nil
}
