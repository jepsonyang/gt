package gtCrypto

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

func Md5(src string) string {
	obj := md5.New()
	obj.Write([]byte(src))
	bytes := obj.Sum(nil)
	return hex.EncodeToString(bytes)
}

func Sha1(src string) string {
	obj := sha1.New()
	obj.Write([]byte(src))
	bytes := obj.Sum(nil)
	return hex.EncodeToString(bytes)
}

func Sha256(src string, key string) string {
	obj := hmac.New(sha256.New, []byte(key))
	obj.Write([]byte(src))
	bytes := obj.Sum(nil)
	return hex.EncodeToString(bytes)
}

func HMACSHA1(src string, key string) string {
	obj := hmac.New(sha1.New, []byte(key))
	obj.Write([]byte(src))
	bytes := obj.Sum(nil)
	return hex.EncodeToString(bytes)
}
