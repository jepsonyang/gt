package gtMd5

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5Sum(src string) string {
	obj := md5.New()
	obj.Write([]byte(src))
	bytes := obj.Sum(nil)
	return hex.EncodeToString(bytes)
}

func Md5SumByte(src []byte) []byte {
	obj := md5.New()
	obj.Write(src)
	return obj.Sum(nil)
}