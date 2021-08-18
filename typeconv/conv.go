// Package to help convert safely interface{} values into specific types
package typeconv

import (
	"fmt"
	"strconv"
)

const IntSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64

func asIntN(n int64, size, bitsize int) (int64, error) {
	// caller is trusted to not lie about size

	if size <= bitsize {
		// same or fewer bits, same signness
		return n, nil
	}

	cutoff := int64(1 << (bitsize - 1))
	if cutoff > n && n >= -cutoff {
		// in range
		return n, nil
	}

	// out of range
	return 0, ErrOutOfRange
}

func asIntN2(u uint64, size, bitsize int) (int64, error) {
	// caller is trusted to not lie about size

	if size < bitsize {
		// fewer bits, despite different signess, always safe.
		return int64(u), nil
	}

	cutoff := uint64(1 << (bitsize - 1))
	if u < cutoff {
		// in range
		return int64(u), nil
	}

	// out of range
	return 0, ErrOutOfRange
}

func asUintN(u uint64, size, bitsize int) (uint64, error) {
	// caller is trusted to not lie about size

	if size <= bitsize {
		// same or fewer bits, same signness
		return u, nil
	}

	cutoff := uint64(1 << bitsize)
	if u < cutoff {
		// in range
		return uint64(u), nil
	}

	// out of range
	return 0, ErrOutOfRange
}

func asUintN2(n int64, size, bitsize int) (uint64, error) {
	// caller is trusted to not lie about size
	if n < 0 {
		// no negatives on uint
		return 0, ErrOutOfRange
	} else if size <= bitsize {
		// same or fewer bits and positive, always safe
		return uint64(n), nil
	}

	cutoff := int64(1 << (bitsize))
	if n < cutoff {
		// in range
		return uint64(n), nil
	}

	// out of range
	return 0, ErrOutOfRange
}

// AsInt tries to convert data into an int
func AsInt(v interface{}) (int, error) {
	if n, ok := v.(int); ok {
		return n, nil
	} else if n, err := AsIntN(v, IntSize); err != nil {
		return 0, err
	} else {
		return int(n), nil
	}
}

// AsInt tries to convert data into a uint
func AsUint(v interface{}) (uint, error) {
	if u, ok := v.(uint); ok {
		return u, nil
	} else if u, err := AsUintN(v, IntSize); err != nil {
		return 0, err
	} else {
		return uint(u), nil
	}
}

// AsInt tries to convert data into a string
func AsString(v interface{}) (string, error) {
	if v == nil {
		return "", InvalidTypeError(v)
	} else if s, ok := v.(string); ok {
		return s, nil
	} else if w, ok := v.(fmt.Stringer); ok {
		return w.String(), nil
	} else {
		s := fmt.Sprintf("%v", v)
		return s, nil
	}
}

// AsBool tries to convert data into a bool
func AsBool(v interface{}) (bool, error) {
	if v == nil {
		return false, InvalidTypeError(v)
	} else if x, ok := v.(bool); ok {
		return x, nil
	} else if s, ok := v.(string); !ok {
		return false, InvalidTypeError(v)
	} else {
		return strconv.ParseBool(s)
	}
}

// AsInt tries to convert data into a slice of something
func AsSlice(v interface{}) ([]interface{}, error) {
	switch w := v.(type) {
	case []interface{}:
		return w, nil
	case interface{}:
		return []interface{}{w}, nil
	default:
		return nil, InvalidTypeError(v)
	}
}
