package conv

import (
	"reflect"
	"unsafe"
)

// https://github.com/gin-gonic/gin/blob/master/internal/bytesconv/bytesconv.go
// StringToBytes converts string to byte slice without a memory allocation.
func StringToBytes(s string) (b []byte) {
	sh := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Len
	return b
}

// https://github.com/gin-gonic/gin/blob/master/internal/bytesconv/bytesconv.go
// BytesToString converts byte slice to string without a memory allocation.
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
