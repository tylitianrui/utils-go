package utils

import "testing"

func TestXrange(t *testing.T) {
	c := Xrange(1, 10, 1)
	for value := range c {
		t.Log(value)

	}
}
