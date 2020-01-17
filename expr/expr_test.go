package expr

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {

	if v, err := New("1 + 3 "); err != nil {
		t.Fail()
	} else {
		res := v.Run()
		fmt.Println(res)
	}
}
