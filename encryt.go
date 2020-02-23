package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
)

// return  md5 checksum value
func MD5(b []byte) string {
	checksum := md5.Sum(b)
	return hex.EncodeToString(checksum[:])

}

// return  sha1 checksum value
func SHA1(b []byte) string {
	checksum := sha1.Sum(b)
	return hex.EncodeToString(checksum[:])

}

// return  sha256 checksum value
func SHA256(b []byte) string {
	checksum := sha256.Sum256(b)
	return hex.EncodeToString(checksum[:])

}

// return  sha512 checksum value
func SHA512(b []byte) string {
	checksum := sha512.Sum512(b)
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
