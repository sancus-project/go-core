// Code generated by ./intsize.sh DO NOT EDIT
package typeconv

//go:generate ./intsize.sh

import (
	"strconv"
)

// AsIntN tries to convert a int64 to int of given size
func AsIntN(v interface{}, bitsize int) (int64, bool) {
	var n int64
	var ok bool

	switch w := v.(type) {
	case string:
		var err error
		n, err = strconv.ParseInt(w, 10, bitsize)
		if err == nil {
			ok = true
		}
	case int:
		n, ok = asIntN(int64(w), IntSize, bitsize)
	case int64:
		n, ok = asIntN(w, 64, bitsize)
	case int32:
		n, ok = asIntN(int64(w), 32, bitsize)
	case int16:
		n, ok = asIntN(int64(w), 16, bitsize)
	case int8:
		n, ok = asIntN(int64(w), 8, bitsize)
	case uint:
		n, ok = asIntN2(uint64(w), IntSize, 64)
	case uint64:
		n, ok = asIntN2(uint64(w), 64, 64)
	case uint32:
		n, ok = asIntN2(uint64(w), 32, 64)
	case uint16:
		n, ok = asIntN2(uint64(w), 16, 64)
	case uint8:
		n, ok = asIntN2(uint64(w), 8, 64)
	}

	return int64(n), ok
}

// AsInt64 tries to convert to int64
func AsInt64(v interface{}) (int64, bool) {
	var n int64
	var ok bool

	switch w := v.(type) {
	case string:
		var err error
		n, err = strconv.ParseInt(w, 10, 64)
		if err == nil {
			ok = true
		}
	case int64:
		return w, true
	case int32:
		return int64(w), true
	case int16:
		return int64(w), true
	case int8:
		return int64(w), true
	case int:
		return int64(w), true
	case uint64:
		n, ok = asIntN2(uint64(w), 64, 64)
	case uint32:
		n, ok = asIntN2(uint64(w), 32, 64)
	case uint16:
		n, ok = asIntN2(uint64(w), 16, 64)
	case uint8:
		n, ok = asIntN2(uint64(w), 8, 64)
	case uint:
		n, ok = asIntN2(uint64(w), IntSize, 64)
	}

	return int64(n), ok
}

// AsInt64Slice tries to convert to slice of int64
func AsInt64Slice(v interface{}) ([]int64, bool) {
	switch w := v.(type) {
	case []int64:
		// ready
		return w, true
	case []interface{}:
		// slice of something else, convert them to Int64
		out := make([]int64, 0, len(w))
		for _, o := range w {
			if n, ok := AsInt64(o); ok {
				out = append(out, n)
			} else {
				// failed
				return nil, false
			}
		}
		return out, true
	case interface{}:
		// promote single element to slice
		if n, ok := AsInt64(v); ok {
			return []int64{n}, true
		}
	}

	return nil, false
}

// AsInt32 tries to convert to int32
func AsInt32(v interface{}) (int32, bool) {
	var n int64
	var ok bool

	switch w := v.(type) {
	case string:
		var err error
		n, err = strconv.ParseInt(w, 10, 32)
		if err == nil {
			ok = true
		}
	case int32:
		return w, true
	case int16:
		return int32(w), true
	case int8:
		return int32(w), true
	case int:
		n, ok = asIntN(int64(w), IntSize, 32)
	case int64:
		n, ok = asIntN(int64(w), 64, 32)
	case uint32:
		n, ok = asIntN2(uint64(w), 32, 32)
	case uint:
		n, ok = asIntN2(uint64(w), IntSize, 32)
	case uint64:
		n, ok = asIntN2(uint64(w), 64, 32)
	case uint16:
		n, ok = asIntN2(uint64(w), 16, 32)
	case uint8:
		n, ok = asIntN2(uint64(w), 8, 32)
	}

	return int32(n), ok
}

// AsInt32Slice tries to convert to slice of int32
func AsInt32Slice(v interface{}) ([]int32, bool) {
	switch w := v.(type) {
	case []int32:
		// ready
		return w, true
	case []interface{}:
		// slice of something else, convert them to Int32
		out := make([]int32, 0, len(w))
		for _, o := range w {
			if n, ok := AsInt32(o); ok {
				out = append(out, n)
			} else {
				// failed
				return nil, false
			}
		}
		return out, true
	case interface{}:
		// promote single element to slice
		if n, ok := AsInt32(v); ok {
			return []int32{n}, true
		}
	}

	return nil, false
}

// AsInt16 tries to convert to int16
func AsInt16(v interface{}) (int16, bool) {
	var n int64
	var ok bool

	switch w := v.(type) {
	case string:
		var err error
		n, err = strconv.ParseInt(w, 10, 16)
		if err == nil {
			ok = true
		}
	case int16:
		return w, true
	case int8:
		return int16(w), true
	case int:
		n, ok = asIntN(int64(w), IntSize, 16)
	case int64:
		n, ok = asIntN(int64(w), 64, 16)
	case int32:
		n, ok = asIntN(int64(w), 32, 16)
	case uint16:
		n, ok = asIntN2(uint64(w), 16, 16)
	case uint:
		n, ok = asIntN2(uint64(w), IntSize, 16)
	case uint64:
		n, ok = asIntN2(uint64(w), 64, 16)
	case uint32:
		n, ok = asIntN2(uint64(w), 32, 16)
	case uint8:
		n, ok = asIntN2(uint64(w), 8, 16)
	}

	return int16(n), ok
}

// AsInt16Slice tries to convert to slice of int16
func AsInt16Slice(v interface{}) ([]int16, bool) {
	switch w := v.(type) {
	case []int16:
		// ready
		return w, true
	case []interface{}:
		// slice of something else, convert them to Int16
		out := make([]int16, 0, len(w))
		for _, o := range w {
			if n, ok := AsInt16(o); ok {
				out = append(out, n)
			} else {
				// failed
				return nil, false
			}
		}
		return out, true
	case interface{}:
		// promote single element to slice
		if n, ok := AsInt16(v); ok {
			return []int16{n}, true
		}
	}

	return nil, false
}

// AsInt8 tries to convert to int8
func AsInt8(v interface{}) (int8, bool) {
	var n int64
	var ok bool

	switch w := v.(type) {
	case string:
		var err error
		n, err = strconv.ParseInt(w, 10, 8)
		if err == nil {
			ok = true
		}
	case int8:
		return w, true
	case int:
		n, ok = asIntN(int64(w), IntSize, 8)
	case int64:
		n, ok = asIntN(int64(w), 64, 8)
	case int32:
		n, ok = asIntN(int64(w), 32, 8)
	case int16:
		n, ok = asIntN(int64(w), 16, 8)
	case uint8:
		n, ok = asIntN2(uint64(w), 8, 8)
	case uint:
		n, ok = asIntN2(uint64(w), IntSize, 8)
	case uint64:
		n, ok = asIntN2(uint64(w), 64, 8)
	case uint32:
		n, ok = asIntN2(uint64(w), 32, 8)
	case uint16:
		n, ok = asIntN2(uint64(w), 16, 8)
	}

	return int8(n), ok
}

// AsInt8Slice tries to convert to slice of int8
func AsInt8Slice(v interface{}) ([]int8, bool) {
	switch w := v.(type) {
	case []int8:
		// ready
		return w, true
	case []interface{}:
		// slice of something else, convert them to Int8
		out := make([]int8, 0, len(w))
		for _, o := range w {
			if n, ok := AsInt8(o); ok {
				out = append(out, n)
			} else {
				// failed
				return nil, false
			}
		}
		return out, true
	case interface{}:
		// promote single element to slice
		if n, ok := AsInt8(v); ok {
			return []int8{n}, true
		}
	}

	return nil, false
}

// AsUintN tries to convert a uint64 to uint of given size
func AsUintN(v interface{}, bitsize int) (uint64, bool) {
	var n uint64
	var ok bool

	switch w := v.(type) {
	case string:
		var err error
		n, err = strconv.ParseUint(w, 10, bitsize)
		if err == nil {
			ok = true
		}
	case uint:
		n, ok = asUintN(uint64(w), IntSize, bitsize)
	case uint64:
		n, ok = asUintN(w, 64, bitsize)
	case uint32:
		n, ok = asUintN(uint64(w), 32, bitsize)
	case uint16:
		n, ok = asUintN(uint64(w), 16, bitsize)
	case uint8:
		n, ok = asUintN(uint64(w), 8, bitsize)
	case int:
		n, ok = asUintN2(int64(w), IntSize, 64)
	case int64:
		n, ok = asUintN2(int64(w), 64, 64)
	case int32:
		n, ok = asUintN2(int64(w), 32, 64)
	case int16:
		n, ok = asUintN2(int64(w), 16, 64)
	case int8:
		n, ok = asUintN2(int64(w), 8, 64)
	}

	return uint64(n), ok
}

// AsUint64 tries to convert to uint64
func AsUint64(v interface{}) (uint64, bool) {
	var n uint64
	var ok bool

	switch w := v.(type) {
	case string:
		var err error
		n, err = strconv.ParseUint(w, 10, 64)
		if err == nil {
			ok = true
		}
	case uint64:
		return w, true
	case uint32:
		return uint64(w), true
	case uint16:
		return uint64(w), true
	case uint8:
		return uint64(w), true
	case uint:
		return uint64(w), true
	case int64:
		n, ok = asUintN2(int64(w), 64, 64)
	case int32:
		n, ok = asUintN2(int64(w), 32, 64)
	case int16:
		n, ok = asUintN2(int64(w), 16, 64)
	case int8:
		n, ok = asUintN2(int64(w), 8, 64)
	case int:
		n, ok = asUintN2(int64(w), IntSize, 64)
	}

	return uint64(n), ok
}

// AsUint64Slice tries to convert to slice of uint64
func AsUint64Slice(v interface{}) ([]uint64, bool) {
	switch w := v.(type) {
	case []uint64:
		// ready
		return w, true
	case []interface{}:
		// slice of something else, convert them to Uint64
		out := make([]uint64, 0, len(w))
		for _, o := range w {
			if n, ok := AsUint64(o); ok {
				out = append(out, n)
			} else {
				// failed
				return nil, false
			}
		}
		return out, true
	case interface{}:
		// promote single element to slice
		if n, ok := AsUint64(v); ok {
			return []uint64{n}, true
		}
	}

	return nil, false
}

// AsUint32 tries to convert to uint32
func AsUint32(v interface{}) (uint32, bool) {
	var n uint64
	var ok bool

	switch w := v.(type) {
	case string:
		var err error
		n, err = strconv.ParseUint(w, 10, 32)
		if err == nil {
			ok = true
		}
	case uint32:
		return w, true
	case uint16:
		return uint32(w), true
	case uint8:
		return uint32(w), true
	case uint:
		n, ok = asUintN(uint64(w), IntSize, 32)
	case uint64:
		n, ok = asUintN(uint64(w), 64, 32)
	case int32:
		n, ok = asUintN2(int64(w), 32, 32)
	case int:
		n, ok = asUintN2(int64(w), IntSize, 32)
	case int64:
		n, ok = asUintN2(int64(w), 64, 32)
	case int16:
		n, ok = asUintN2(int64(w), 16, 32)
	case int8:
		n, ok = asUintN2(int64(w), 8, 32)
	}

	return uint32(n), ok
}

// AsUint32Slice tries to convert to slice of uint32
func AsUint32Slice(v interface{}) ([]uint32, bool) {
	switch w := v.(type) {
	case []uint32:
		// ready
		return w, true
	case []interface{}:
		// slice of something else, convert them to Uint32
		out := make([]uint32, 0, len(w))
		for _, o := range w {
			if n, ok := AsUint32(o); ok {
				out = append(out, n)
			} else {
				// failed
				return nil, false
			}
		}
		return out, true
	case interface{}:
		// promote single element to slice
		if n, ok := AsUint32(v); ok {
			return []uint32{n}, true
		}
	}

	return nil, false
}

// AsUint16 tries to convert to uint16
func AsUint16(v interface{}) (uint16, bool) {
	var n uint64
	var ok bool

	switch w := v.(type) {
	case string:
		var err error
		n, err = strconv.ParseUint(w, 10, 16)
		if err == nil {
			ok = true
		}
	case uint16:
		return w, true
	case uint8:
		return uint16(w), true
	case uint:
		n, ok = asUintN(uint64(w), IntSize, 16)
	case uint64:
		n, ok = asUintN(uint64(w), 64, 16)
	case uint32:
		n, ok = asUintN(uint64(w), 32, 16)
	case int16:
		n, ok = asUintN2(int64(w), 16, 16)
	case int:
		n, ok = asUintN2(int64(w), IntSize, 16)
	case int64:
		n, ok = asUintN2(int64(w), 64, 16)
	case int32:
		n, ok = asUintN2(int64(w), 32, 16)
	case int8:
		n, ok = asUintN2(int64(w), 8, 16)
	}

	return uint16(n), ok
}

// AsUint16Slice tries to convert to slice of uint16
func AsUint16Slice(v interface{}) ([]uint16, bool) {
	switch w := v.(type) {
	case []uint16:
		// ready
		return w, true
	case []interface{}:
		// slice of something else, convert them to Uint16
		out := make([]uint16, 0, len(w))
		for _, o := range w {
			if n, ok := AsUint16(o); ok {
				out = append(out, n)
			} else {
				// failed
				return nil, false
			}
		}
		return out, true
	case interface{}:
		// promote single element to slice
		if n, ok := AsUint16(v); ok {
			return []uint16{n}, true
		}
	}

	return nil, false
}

// AsUint8 tries to convert to uint8
func AsUint8(v interface{}) (uint8, bool) {
	var n uint64
	var ok bool

	switch w := v.(type) {
	case string:
		var err error
		n, err = strconv.ParseUint(w, 10, 8)
		if err == nil {
			ok = true
		}
	case uint8:
		return w, true
	case uint:
		n, ok = asUintN(uint64(w), IntSize, 8)
	case uint64:
		n, ok = asUintN(uint64(w), 64, 8)
	case uint32:
		n, ok = asUintN(uint64(w), 32, 8)
	case uint16:
		n, ok = asUintN(uint64(w), 16, 8)
	case int8:
		n, ok = asUintN2(int64(w), 8, 8)
	case int:
		n, ok = asUintN2(int64(w), IntSize, 8)
	case int64:
		n, ok = asUintN2(int64(w), 64, 8)
	case int32:
		n, ok = asUintN2(int64(w), 32, 8)
	case int16:
		n, ok = asUintN2(int64(w), 16, 8)
	}

	return uint8(n), ok
}

// AsUint8Slice tries to convert to slice of uint8
func AsUint8Slice(v interface{}) ([]uint8, bool) {
	switch w := v.(type) {
	case []uint8:
		// ready
		return w, true
	case []interface{}:
		// slice of something else, convert them to Uint8
		out := make([]uint8, 0, len(w))
		for _, o := range w {
			if n, ok := AsUint8(o); ok {
				out = append(out, n)
			} else {
				// failed
				return nil, false
			}
		}
		return out, true
	case interface{}:
		// promote single element to slice
		if n, ok := AsUint8(v); ok {
			return []uint8{n}, true
		}
	}

	return nil, false
}
