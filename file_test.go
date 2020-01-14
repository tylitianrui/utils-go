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
