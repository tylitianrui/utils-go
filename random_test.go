package utils

import "testing"

func TestRandInt(t *testing.T) {
	for i := 0; i < 100; i++ {
		r, _ := RandInt(-12, 12)
		if r >= -12 && r < 12 {
			t.Log(r)
		} else {
			t.Fail()
		}

	}

}
