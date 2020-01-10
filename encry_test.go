package utils

import "testing"

func TestMD5(t *testing.T) {
	t.Log(MD5([]byte("123")))
}
