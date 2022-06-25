package errors

import (
	"strings"
)

type Validator interface {
	Ok() bool
	Error() string
	Errors() []error
}

func AsValidator(err error) (Validator, bool) {
	for err != nil {
		if v, ok := err.(Validator); ok {
			return v, true
		}

		err = Unwrap(err)
	}
	return nil, false
}

type ErrorStack struct {
	errors []error
}

func (s ErrorStack) Ok() bool {
	return len(s.errors) == 0
}

func (s ErrorStack) Error() string {
	var errors []string

	for _, err := range s.errors {
		errors = append(errors, err.Error())
	}

	return strings.Join(errors, "\n")
}

// Errors() returns the included errors
func (s ErrorStack) Errors() []error {
	return s.errors
}

// AsError() returns nil of Ok() or itself if not
func (s *ErrorStack) AsError() error {
	if !s.Ok() {
		return s
	}
	return nil
}

// Appends an error
func (s *ErrorStack) AppendError(err error) {
	if err == nil {
		// skip
	} else if nerr, ok := AsValidator(err); !ok {
		// simple
		s.errors = append(s.errors, err)
	} else if !nerr.Ok() {
		// nested
		s.errors = append(s.errors, nerr.Errors()...)
	}
}

// Appends a described error
func (s *ErrorStack) AppendErrorf(str string, args ...interface{}) {
	if len(str) > 0 {
		s.AppendError(New(str, args...))
	}
}

// Appends a described error wrapping another error
func (s *ErrorStack) AppendWrapped(err error, str string, args ...interface{}) {
	if len(str) > 0 {
		err = Wrap(err, str, args...)
	}
	s.AppendError(err)
}

// New ErrorStack wrapping a given list of errors
func NewErrorStack(errs ...error) ErrorStack {

	j := 0
	for i, err := range errs {
		if err == nil {
			continue
		}

		if i != j {
			errs[j] = err
		}
		j++
	}

	return ErrorStack{
		errors: errs[:j],
	}
}
