package utilc

import (
	"crypto/md5"
	"encoding/hex"
)

func If(exp bool, a interface{}, b interface{}) interface{} {
	if exp {
		return a
	}
	return b
}

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
