package errors

import (
	"fmt"
)

type StackTracer interface {
	StackTrace() Stack
}

type StackError struct {
	stack Stack
	err   error
}

func (s *StackError) Error() string {
	if s.err != nil {
		return s.err.Error()
	} else {
		return "stacktrace"
	}
}

func (s *StackError) StackTrace() Stack {
	return s.stack
}

func (s *StackError) Unwrap() error {
	return s.err
}

func WithStackTrace(skip int, err error) error {
	return &StackError{
		err:   err,
		stack: BackTrace(1 + skip),
	}
}

type Panic interface {
	Recovered() interface{}
}

type PanicError struct {
	stack Stack
	rvr   interface{}
}

func (p *PanicError) Error() string {
	return fmt.Sprintf("panic: %s", p.rvr)
}

func (p *PanicError) StackTrace() Stack {
	return p.stack
}

func (p *PanicError) Unwrap() error {
	if err, ok := p.rvr.(error); ok {
		return err
	} else {
		return nil
	}
}

func (p *PanicError) Recovered() interface{} {
	return p.rvr
}

func Recover() *PanicError {
	if rvr := recover(); rvr == nil {
		// no error
		return nil
	} else if p, ok := rvr.(*PanicError); ok {
		// pass previous panic along
		return p
	} else {
		// spawn new PanicError
		return panicError(rvr)
	}
}

func panicError(rvr interface{}) *PanicError {
	return &PanicError{
		rvr:   rvr,
		stack: BackTrace(2),
	}
}
