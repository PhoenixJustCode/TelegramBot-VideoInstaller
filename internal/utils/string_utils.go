package utils

import "strings"

func IsYoutubeURL(url string) bool {
    return strings.Contains(url, "youtube") || strings.Contains(url, "youtu.be")
}


func IsMP4File(url string) bool {
    return strings.HasSuffix(url, ".mp4") 
}