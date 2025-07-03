package utils

import "strings"

func IsVideoURL(url string) bool {
    return strings.HasSuffix(url, ".mp4") || strings.Contains(url, "youtube")
}
