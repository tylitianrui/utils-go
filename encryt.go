package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// return  md5 checksum value
func MD5(b []byte) string {
	checksum := md5.Sum(b)
	return hex.EncodeToString(checksum[:])

}
