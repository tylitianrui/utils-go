package utils

import "time"

// 时间戳，秒级
func Timestamp() int64 {
	return time.Now().Unix()
}

// 时间戳，纳秒级
func TimestampNano() int64 {
	return time.Now().UnixNano()
}

// 时间戳，微级
func TimestampMicro() int64 {
	return time.Now().UnixNano() / 1000
}

// 时间戳，毫级
func TimestampMilli() int64 {
	return time.Now().UnixNano() / 1000000
}
