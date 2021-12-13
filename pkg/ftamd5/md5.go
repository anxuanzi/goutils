package ftamd5

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5Hash 返回32字节小写字符串
func Md5Hash(b []byte) string {
	h := md5.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}
