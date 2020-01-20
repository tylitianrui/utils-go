package utils

type ErrLV int

type MyError struct {
	msg   string
	level ErrLV
}

func NewError(msg string, level ErrLV) *MyError {
	return &MyError{
		msg:   msg,
		level: level,
	}
}

func (self *MyError) Error() string {
	return self.msg
}
