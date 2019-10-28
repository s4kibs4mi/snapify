package utils

import "strings"

func FormatUrlWithoutProtocol(url string) string {
	nUrl := strings.TrimPrefix(url, "https://")
	if url != nUrl {
		return nUrl
	}
	return strings.TrimPrefix(url, "http://")
}

func FormatUrlWithProtocol(url string) string {
	if strings.HasPrefix(url, "https://") || strings.HasPrefix(url, "http://") {
		return url
	}
	return "http://" + url
}
