package utils

import (
	"errors"
	"math/rand"
	"time"
)

// a random  number in [start,end)
// if  start >= end ,return  error
func RandInt(start, end int) (int, error) {
	if start >= end {
		return 0, errors.New("start should less than end")
	}
	seed := time.Now().UnixNano()
	r := end - start
	rand.Seed(seed)
	return rand.Intn(r) + start, nil
}
