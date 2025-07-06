package utils

import (
	"errors"
	"regexp"
	"strings"
)

func ExtractVideoID(url string) (string, error) {
	if strings.Contains(url, "search_query") || strings.Contains(url, "/channel/") || strings.Contains(url, "/user/") {
		return "", errors.New("❌ Это не ссылка на видео YouTube")
	}

	re := regexp.MustCompile(`(?:v=|\/embed\/|youtu\.be\/)([a-zA-Z0-9_-]{11})`)
	match := re.FindStringSubmatch(url)
	if len(match) < 2 {
		return "", errors.New("❌ Не удалось извлечь ID видео")
	}

	return match[1], nil
}
