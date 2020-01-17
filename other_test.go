package utils

import "testing"

func TestConvert10toN(t *testing.T) {
	res := Convert10toN(64, 62)
	if res == "12" {
		t.Log(res)
	} else {
		t.Fail()

	}

	res1 := Convert10toN(64, 12)
	if res1 == "54" {
		t.Log(res1)
	} else {
		t.Fail()

	}

	res2 := Convert10toN(64, 2)
	t.Log(res2)
	if res2 == "54" {
		t.Log(res1)
	} else {
		t.Fail()

	}
}
