package pretty

import (
	"github.com/tylitianrui/utils-go/conv"
	"strconv"
	"time"
)

const capacity = 1 << 10

type Buffer struct {
	b []byte
}

func (b *Buffer) Write(p []byte) (n int, err error) {
	b.b = append(b.b, p...)
	return len(p), nil
}

func (b *Buffer) String() string {
	return conv.BytesToString(b.b)
}

func (b *Buffer) Append(buf *Buffer) (n int, err error) {
	return b.Write(buf.b)
}

func (b *Buffer) AppendString(s ...string) {
	for i := 0; i < len(s); i++ {
		b.b = append(b.b, s[i]...)
	}
}

func (b *Buffer) AppendBytes(p ...byte) (n int, err error) {
	return b.Write(p)
}

func (b *Buffer) AppendTime(t time.Time, layout string) {
	b.b = t.AppendFormat(b.b, layout)
}

func (b *Buffer) AppendInt64(i int64) {
	b.b = strconv.AppendInt(b.b, i, 10)
}

func (b *Buffer) AppendUint64(i uint64) {
	b.b = strconv.AppendUint(b.b, i, 10)
}

func (b *Buffer) AppendFloat64(i float64, bitSize int) {
	b.b = strconv.AppendFloat(b.b, i, 'f', -1, bitSize)
}

func (b *Buffer) Reset() {
	b.b = b.b[:0]
}

func newBuffer() *Buffer {
	return &Buffer{make([]byte, capacity)}
}
