package errors

import (
	"fmt"
)

type MissingFieldError struct {
	field string
}

func (e MissingFieldError) Error() string {
	if len(e.field) > 0 {
		return fmt.Sprintf("%s: %s", "Missing Field", e.field)
	}
	return "Missing Field"
}

func ErrMissingField(str string, args ...interface{}) *MissingFieldError {
	if len(str) == 0 {
		return nil
	}

	if len(args) > 0 {
		str = fmt.Sprintf(str, args...)
	}

	return &MissingFieldError{
		field: str,
	}
}

type MissingArgumentError struct {
	arg string
}

func (e MissingArgumentError) Error() string {
	if len(e.arg) > 0 {
		return fmt.Sprintf("%s: %s", "Missing Argument", e.arg)
	}
	return "Missing Argument"
}

func ErrMissingArgument(str string, args ...interface{}) *MissingArgumentError {
	if len(str) == 0 {
		return nil
	}

	if len(args) > 0 {
		str = fmt.Sprintf(str, args...)
	}

	return &MissingArgumentError{
		arg: str,
	}
}
