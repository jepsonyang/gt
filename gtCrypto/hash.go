package gtCrypto

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
)

func Md5(src string) []byte {
	obj := md5.New()
	obj.Write([]byte(src))
	return obj.Sum(nil)
}

func Sha1(src string) []byte {
	obj := sha1.New()
	obj.Write([]byte(src))
	return obj.Sum(nil)
}

func Sha256(src string, key string) []byte {
	obj := hmac.New(sha256.New, []byte(key))
	obj.Write([]byte(src))
	return obj.Sum(nil)
}

func HMACSha1(src string, key string) []byte {
	obj := hmac.New(sha1.New, []byte(key))
	obj.Write([]byte(src))
	return obj.Sum(nil)
}
