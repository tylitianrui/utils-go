package utils

import "testing"

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
