package utils

import "testing"

func TestRandLetter(t *testing.T) {
	t1 := RandLetter(12, LetterAll)
	t.Log(t1)
}
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

//func TestCreateRandomString(t *testing.T) {
//	for i:=0; i < 100; i++ {
//		l:= RandUpperLetter(12)
//		if  len(l)==12{
//			t.Log(l)
//		}else {
//			t.Fail()
//		}
//	}
//
//}
