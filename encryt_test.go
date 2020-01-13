package utils

import "testing"

func TestMD5(t *testing.T) {
	t.Log(MD5([]byte("123")))
}

func TestBase64(t *testing.T) {
	s := []byte("hell0")
	e := Base64(s)
	t.Logf("content :%s base64:%s", s, e)

	if c, err := DecodeBase64(e); err != nil {
		t.Fail()
	} else {
		t.Log(string(c))

	}
}
