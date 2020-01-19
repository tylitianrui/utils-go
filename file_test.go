package utils

import (
	"log"
	"testing"
)

func TestCurrentPath(t *testing.T) {
	t.Log(CurrentPath())
}

func TestSubFileName(t *testing.T) {
	l, err := SubFileName("./")
	if err != nil {
		log.Println(err)
	} else {
		log.Println(len(l))
	}
}

func TestCopy(t *testing.T) {
	src := "/Users/tyltr/opencode/utils-go/tyltr/mian.go"
	dst := "/Users/tyltr/opencode/utils-go/tyltr/1/q.go"
	if err := copy(src, dst); err != nil {
		t.Log(err)
		t.Fail()
	}
}

func TestCopyFile(t *testing.T) {
	s := "/Users/tyltr/opencode/utils-go/tyltr/mian.go"
	d := "/Users/tyltr/opencode/utils-go/tyltr/1/q.go"

	if err := CopyFile(s, d, "/1"); err != nil {
		t.Log(err)
		t.Fail()
	}
}
