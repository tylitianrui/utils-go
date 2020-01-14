package set

import (
	"log"
	"testing"
)

func TestHashSet_Add(t *testing.T) {
	set := NewHashSet()
	set.Add(1)
	set.Add(1)
	set.Add(1)
	set.Add(1)
	set.Add("q")
	if set.Len() == 2 {
		log.Println("ok")
	} else {
		t.Fail()
	}

	set.Pop(1)
	if set.Len() == 1 {
		log.Println("ok")
	} else {
		t.Fail()
	}

}
