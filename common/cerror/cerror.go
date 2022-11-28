package cerror

import "fmt"

var (
	ErrUnKnownError = func(err error) error {
		return New("E002", err.Error())
	}
)

type Error struct {
	code string
	msg  string
}

func New(code string, message string) error {
	return &Error{
		code: code,
		msg:  message,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error Code: %s, Error: %s", e.code, e.msg)
}

func (e *Error) Code() string {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}
