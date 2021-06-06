package errors

import (
	"fmt"
)

type Panic interface {
	StackTracer

	Error() string
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

func Recover() Panic {
	if rvr := recover(); rvr == nil {
		// no error
		return nil
	} else if p, ok := rvr.(Panic); ok {
		// pass previous panic along
		return p
	} else {
		// spawn new PanicError
		return panicError(rvr)
	}
}

func panicError(rvr interface{}) Panic {
	return &PanicError{
		rvr:   rvr,
		stack: BackTrace(2),
	}
}
