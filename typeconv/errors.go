package typeconv

import (
	"go.sancus.dev/core/errors"
)

var (
	ErrOutOfRange = errors.New("Value out of range")
)

func InvalidTypeError(v interface{}) error {
	return errors.New("Invalid Type: %T", v)
}
