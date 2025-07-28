package core

import (
	"fmt"
)

type appErr interface {
	Code() int
	Error() string
	Update(context string) appErr
}

type appError struct {
	code  int
	error string
}

func (a appError) Code() int {
	return a.code
}

func (a appError) Error() string {
	return fmt.Sprintf("Code %d: %s", a.code, a.Error)
}

func (a appError) Update(context string) appErr {
	return appError{code: a.code, error: fmt.Sprintf("when: %s: %s", context, a.error)}
}

func MakeAppErr(code int, context string, err error) appError {
	return appError{code: code, error: fmt.Sprintf("when: %s: %e", context, err)}
}
