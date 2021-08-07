// Package to help convert safely interface{} values into specific types
package typeconv

import (
	"fmt"
)

const IntSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64

func asIntN(n int64, size, bitsize int) (int64, bool) {
	// caller is trusted to not lie about size

	if size <= bitsize {
		// same or fewer bits, same signness
		return n, true
	}

	cutoff := int64(1 << (bitsize - 1))
	if cutoff > n && n >= -cutoff {
		// in range
		return n, true
	}

	// out of range
	return 0, false
}

func asIntN2(u uint64, size, bitsize int) (int64, bool) {
	// caller is trusted to not lie about size

	if size < bitsize {
		// fewer bits, despite different signess, always safe.
		return int64(u), true
	}

	cutoff := uint64(1 << (bitsize - 1))
	if u < cutoff {
		// in range
		return int64(u), true
	}

	// out of range
	return 0, false
}

func asUintN(u uint64, size, bitsize int) (uint64, bool) {
	// caller is trusted to not lie about size

	if size <= bitsize {
		// same or fewer bits, same signness
		return u, true
	}

	cutoff := uint64(1 << bitsize)
	if u < cutoff {
		// in range
		return uint64(u), true
	}

	// out of range
	return 0, false
}

func asUintN2(n int64, size, bitsize int) (uint64, bool) {
	// caller is trusted to not lie about size
	if n < 0 {
		// no negatives on uint
		return 0, false
	} else if size <= bitsize {
		// same or fewer bits and positive, always safe
		return uint64(n), true
	}

	cutoff := int64(1 << (bitsize))
	if n < cutoff {
		// in range
		return uint64(n), true
	}

	// out of range
	return 0, false
}

// AsInt tries to convert data into an int
func AsInt(v interface{}) (int, bool) {
	if n, ok := v.(int); ok {
		return n, true
	} else if n, ok := AsIntN(v, IntSize); ok {
		return int(n), true
	} else {
		return 0, false
	}
}

// AsInt tries to convert data into a uint
func AsUint(v interface{}) (uint, bool) {
	if u, ok := v.(uint); ok {
		return u, true
	} else if u, ok := AsUintN(v, IntSize); ok {
		return uint(u), true
	} else {
		return 0, false
	}
}

// AsInt tries to convert data into a string
func AsString(v interface{}) (string, bool) {
	if v == nil {
		return "", false
	} else if s, ok := v.(string); ok {
		return s, true
	} else if w, ok := v.(fmt.Stringer); ok {
		return w.String(), true
	} else {
		s := fmt.Sprintf("%v", v)
		return s, true
	}
}

// AsInt tries to convert data into a slice of something
func AsSlice(v interface{}) ([]interface{}, bool) {
	switch w := v.(type) {
	case []interface{}:
		return w, true
	case interface{}:
		return []interface{}{w}, true
	default:
		return nil, false
	}
}
