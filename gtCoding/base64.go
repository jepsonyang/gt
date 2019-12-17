package gtCoding

import "encoding/base64"

func Base64Encode(src string) string {
	return base64.StdEncoding.EncodeToString([]byte(src))
}

func Base64Decode(src string) (string, error) {
	retByte, err := base64.StdEncoding.DecodeString(src)
	return string(retByte), err
}
