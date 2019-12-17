package gtNet

import "net/url"

func UrlEncode(src string) string {
	return url.QueryEscape(src)
}

func UrlDecode(src string) (string, error) {
	return url.QueryUnescape(src)
}