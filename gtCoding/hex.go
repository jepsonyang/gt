package gtCoding

import (
	"encoding/hex"
)

func HexEncode(src []byte) string {
	return hex.EncodeToString(src)
}

func HexDecode(src string) ([]byte, error) {
	return hex.DecodeString(src)
}
