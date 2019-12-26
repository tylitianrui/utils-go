package utils

import (
	"testing"
)

type S1 struct {
	A int
	B string
}

type S2 struct {
	A S1
	B string
}

func TestGetStructField(t *testing.T) {
	var (
		s1 = S1{}
		s2 S2
	)
	res1 := GetStructField(s1, "A", 12)

	if res1 == 12 {
		t.Log(res1)
	} else {
		t.Fail()
	}

	res2 := GetStructField(s2, "A", 12)
	if res2 == 12 {
		t.Log(res2)
	} else {
		t.Fail()
	}
}
