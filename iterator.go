package utils

const iterLen = 1 << 4

// python中的xrange函数
func Xrange(start, end, step int) chan int {
	c := make(chan int, iterLen)

	go func() {
		for i := start; i < end; i += step {
			c <- i
		}
		close(c)
	}()
	return c
}
