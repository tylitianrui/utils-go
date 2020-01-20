package utils

import (
	"testing"
	"time"
)

func TestTimestamp(t *testing.T) {
	t.Logf("current timestamp(secodes):%d", Timestamp())
}

func TestTimestampNano(t *testing.T) {
	t.Logf("current timestamp(nano-secodes):%d", TimestampNano())
}

func TestTimestampMicro(t *testing.T) {
	t.Logf("current timestamp(micro-secodes):%d", TimestampMicro())
}

func TestTimestampMilli(t *testing.T) {
	t.Logf("current timestamp(milli-secodes):%d", TimestampMilli())

}

func TestTime2Stamp(t *testing.T) {
	//now := time.Now()
	timeTemplate := "2006-01-02 15:04:05"
	t1 := "2019-01-08 13:50:30"
	//time,_:=time.ParseInLocation(timeTemplate, t1, time.Local)
	time, _ := time.Parse(timeTemplate, t1)
	t.Log(time)
}

func TestAddTime(t *testing.T) {
	h, e := AddTime(time.Now(), "-1d")
	t.Log(e)
	t.Log(h)
}
