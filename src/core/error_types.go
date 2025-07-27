package core

import (
	"fmt"
)

type appError interface {
	Code() int
	Error() string
}

type aError struct {
	code  int
	error string
}

func (a aError) Code() int {
	return a.code
}

func (a aError) Error() string {
	return fmt.Sprintf("Code %d: %s", a.code, a.Error)
}
