package errx

import (
	"fmt"
)

type Error interface {
	error
	Cause() error
	WithCause(err error) Error
	WithOther(err ...error) Error
}

type simpleError struct {
	file   string
	line   int
	msg    string
	cause  error
	others []error
}

func (e *simpleError) Error() string {
	s := fmt.Sprintf("%s @ %s:%d", e.msg, e.file, e.line)

	if e.cause != nil {
		s += " caused by " + e.cause.Error()
	}

	if len(e.others) > 0 {
		for _, oe := range e.others {
			s += " with " + oe.Error()
		}
	}

	return s
}

func (e *simpleError) Cause() error {
	return e.cause
}

func (e *simpleError) WithCause(err error) Error {
	e.cause = err
	return e
}

func (e *simpleError) WithOther(err ...error) Error {
	e.others = err
	return e
}

func New(msg string) Error {
	file, line := getCallerInfo(2)
	return &simpleError{file: file, line: line, msg: msg}
}

func Newf(msg string, args ...interface{}) Error {
	return New(fmt.Sprintf(msg, args...))
}

func Cause(err error) error {
	for {
		if e, ok := err.(Error); ok {
			err = e.Cause()
		} else {
			return err
		}
	}
}
