package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

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

var day = regexp.MustCompile(`^(-?\d+d).*`)

// Valid d are "ns", "us" (or "µs"), "ms", "s", "m", "h" ,"d"
func AddTime(base time.Time, d string) (time.Time, error) {
	var (
		dur1, dur time.Duration
		err       error
	)
	v := day.FindStringSubmatch(d)
	if len(v) > 2 {
		return base, errors.New(fmt.Sprintf("time: unknown unit d in duration %s", d))
	} else if len(v) == 2 {
		t := v[1]
		day := t[:len(t)-1]
		d = d[len(day):]
		di, _ := strconv.Atoi(day)
		if dur1, err = time.ParseDuration(fmt.Sprintf("%dh", di*24)); err != nil {
			return base, err
		}
	}

	if v[0] == v[1] {
		dur = time.Duration(0)
	} else if dur, err = time.ParseDuration(d); err != nil {
		return base, err
	}
	return base.Add(dur + dur1), nil
}
