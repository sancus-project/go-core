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

func (s ErrorStack) Errors() []error {
	return s.errors
}

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

func (s *ErrorStack) AppendErrorf(str string, args ...interface{}) {
	if len(str) > 0 {
		s.AppendError(New(str, args...))
	}
}

func (s *ErrorStack) MissingField(str string, args ...interface{}) {
	if err := ErrMissingField(str, args...); err != nil {
		s.errors = append(s.errors, err)
	}
}

func (s *ErrorStack) MissingArgument(str string, args ...interface{}) {
	if err := ErrMissingArgument(str, args...); err != nil {
		s.errors = append(s.errors, err)
	}
}
