package util

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(str string) string {
	h := md5.New()
	var bytes = []byte(str)
	h.Write(bytes)
	return hex.EncodeToString(h.Sum(nil))
}
