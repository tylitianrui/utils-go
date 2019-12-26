package utils

import (
	"errors"
	"math/rand"
	"time"
)

type RandKind int

const (
	LetterAll RandKind = iota
	LetterUpper
	LetterLower
)
const (
	NumLetter  = 26
	UpperStart = 65
	LowerStart = 97
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// a random  number in [start,end)
// if  start >= end ,return  error
func RandInt(start, end int) (int, error) {
	if start >= end {
		return 0, errors.New("start should less than end")
	}
	r := end - start
	return rand.Intn(r) + start, nil
}

//
func RandLetter(len int, kind RandKind) string {
	b := make([]byte, len)
	var (
		s int
		l int
	)
	switch kind {
	case LetterLower:
		s = LowerStart
		l = NumLetter

	case UpperStart:
		s = UpperStart
		l = NumLetter
	default:
		s = UpperStart
		l = NumLetter * 2

	}
	for i := 0; i < len; i++ {
		idx := rand.Intn(l)

		if idx >= NumLetter {
			idx += 7
		}

		b[i] = byte(idx + s)
	}

	return string(b)
}
