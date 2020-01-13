package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

// return  md5 checksum value
func MD5(b []byte) string {
	checksum := md5.Sum(b)
	return hex.EncodeToString(checksum[:])

}

// 获取base64
func Base64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

// 解析base64
func DecodeBase64(b string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(b)
}
