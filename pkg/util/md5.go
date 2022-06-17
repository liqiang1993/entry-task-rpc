package util

import (
	"crypto/md5" //nolint:gosec
	"encoding/hex"
)

// EncodeMD5 md5 encryption
func EncodeMD5(value string) string {
	m := md5.New() //nolint:gosec
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}
