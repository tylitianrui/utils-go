package pretty

import "sync"

var prettyPool = &sync.Pool{
	New: func() interface{} {
		return newBuffer()
	},
}

//Buffer
func Get() *Buffer {
	buf := prettyPool.Get().(*Buffer)
	buf.Reset()
	return buf
}

func Put(buf *Buffer) {
	prettyPool.Put(buf)
}
