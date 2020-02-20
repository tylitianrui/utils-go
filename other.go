package utils

import (
	"fmt"
	"os"
	"regexp"
)

var (
	NumBase = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
)

func Convert10toN(input, n int) string {
	if n <= 1 {
		panic("param  error")
	}
	if n > len(NumBase) {
		panic("超出表示范围，请修改NumBase或者参数n")
	}
	if input < n {
		return fmt.Sprintf("%s", NumBase[input])
	}
	return fmt.Sprintf("%s%s", Convert10toN(input/n, n), NumBase[input%n])
}

var emailRegexp = regexp.MustCompile(`^[\da-z]+([\-\.\_]?[\da-z]+)*@[\da-z]+([\-\.]?[\da-z]+)*(\.[a-z]{2,})+$`)

func VerifyEmail(email *string) bool {
	s := emailRegexp.FindString(*email)
	if s != "" {
		return true
	}
	return false
}

// 获取环境变量 key的值，不存在则返回默认值defval
func GetEnvWithDefault(key, defval string) string {
	var (
		ok  bool
		val string
	)
	if val, ok = os.LookupEnv(key); ok {
		return val
	}
	return defval

}
