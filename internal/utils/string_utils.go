package utils

import "strings"

func IsYoutubeURL(url string) bool {
	return (strings.Contains(url, "youtube.com/watch?v=") || strings.Contains(url, "youtu.be/")) &&
		!strings.Contains(url, "search_query") &&
		!strings.Contains(url, "/channel/") &&
		!strings.Contains(url, "/user/")
}

func IsMP4File(url string) bool {
	return strings.HasSuffix(url, ".mp4")
}
