package lib

import (
	"fmt"
	"strings"
	"time"
)

func GenerateSlug(name string) string {
	return strings.ToLower(strings.ReplaceAll(name, " ", "-"))
}

func GenerateUniqueSlug(slug string) string {
	return fmt.Sprintf("%s-%d", slug, time.Now().UnixNano())
}
