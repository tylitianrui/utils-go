package utils

import "time"

const TimeTemplate = "2019-01-08 13:50:30"

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

// 时间转时间戳，纳秒级
func Time2StampNano(time time.Time) int64 {
	return time.UnixNano()

}

// 时间转时间戳，秒级
func Time2Stamp(time time.Time) int64 {
	return time.Unix()

}

// 字符串解析时间
func TimeParse(timestr, tmplate string) (time.Time, error) {
	return time.Parse(tmplate, timestr)

}
