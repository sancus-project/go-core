package typeconv

//go:generate ./slice.sh String Bool Int Int64 Int32 Int16 Int8 Uint Uint64 Uint32 Uint16 Uint8

// Code generated by ./slice.sh DO NOT EDIT

// AsStringSlice tries to convert data into a slice of string
func AsStringSlice(v interface{}) ([]string, error) {
	switch w := v.(type) {
	case []string:
		// ready
		return w, nil
	case []interface{}:
		// slice of something else, convert them to String
		out := make([]string, 0, len(w))
		for _, o := range w {
			if n, err := AsString(o); err == nil {
				out = append(out, n)
			} else {
				// failed
				return nil, err
			}
		}
		return out, nil
	case interface{}:
		// promote single element to slice
		if n, err := AsString(v); err != nil {
			return nil, err
		} else {
			return []string{n}, nil
		}
	default:
		return nil, InvalidTypeError(v)
	}
}

// AsBoolSlice tries to convert data into a slice of bool
func AsBoolSlice(v interface{}) ([]bool, error) {
	switch w := v.(type) {
	case []bool:
		// ready
		return w, nil
	case []interface{}:
		// slice of something else, convert them to Bool
		out := make([]bool, 0, len(w))
		for _, o := range w {
			if n, err := AsBool(o); err == nil {
				out = append(out, n)
			} else {
				// failed
				return nil, err
			}
		}
		return out, nil
	case interface{}:
		// promote single element to slice
		if n, err := AsBool(v); err != nil {
			return nil, err
		} else {
			return []bool{n}, nil
		}
	default:
		return nil, InvalidTypeError(v)
	}
}

// AsIntSlice tries to convert data into a slice of int
func AsIntSlice(v interface{}) ([]int, error) {
	switch w := v.(type) {
	case []int:
		// ready
		return w, nil
	case []interface{}:
		// slice of something else, convert them to Int
		out := make([]int, 0, len(w))
		for _, o := range w {
			if n, err := AsInt(o); err == nil {
				out = append(out, n)
			} else {
				// failed
				return nil, err
			}
		}
		return out, nil
	case interface{}:
		// promote single element to slice
		if n, err := AsInt(v); err != nil {
			return nil, err
		} else {
			return []int{n}, nil
		}
	default:
		return nil, InvalidTypeError(v)
	}
}

// AsInt64Slice tries to convert data into a slice of int64
func AsInt64Slice(v interface{}) ([]int64, error) {
	switch w := v.(type) {
	case []int64:
		// ready
		return w, nil
	case []interface{}:
		// slice of something else, convert them to Int64
		out := make([]int64, 0, len(w))
		for _, o := range w {
			if n, err := AsInt64(o); err == nil {
				out = append(out, n)
			} else {
				// failed
				return nil, err
			}
		}
		return out, nil
	case interface{}:
		// promote single element to slice
		if n, err := AsInt64(v); err != nil {
			return nil, err
		} else {
			return []int64{n}, nil
		}
	default:
		return nil, InvalidTypeError(v)
	}
}

// AsInt32Slice tries to convert data into a slice of int32
func AsInt32Slice(v interface{}) ([]int32, error) {
	switch w := v.(type) {
	case []int32:
		// ready
		return w, nil
	case []interface{}:
		// slice of something else, convert them to Int32
		out := make([]int32, 0, len(w))
		for _, o := range w {
			if n, err := AsInt32(o); err == nil {
				out = append(out, n)
			} else {
				// failed
				return nil, err
			}
		}
		return out, nil
	case interface{}:
		// promote single element to slice
		if n, err := AsInt32(v); err != nil {
			return nil, err
		} else {
			return []int32{n}, nil
		}
	default:
		return nil, InvalidTypeError(v)
	}
}

// AsInt16Slice tries to convert data into a slice of int16
func AsInt16Slice(v interface{}) ([]int16, error) {
	switch w := v.(type) {
	case []int16:
		// ready
		return w, nil
	case []interface{}:
		// slice of something else, convert them to Int16
		out := make([]int16, 0, len(w))
		for _, o := range w {
			if n, err := AsInt16(o); err == nil {
				out = append(out, n)
			} else {
				// failed
				return nil, err
			}
		}
		return out, nil
	case interface{}:
		// promote single element to slice
		if n, err := AsInt16(v); err != nil {
			return nil, err
		} else {
			return []int16{n}, nil
		}
	default:
		return nil, InvalidTypeError(v)
	}
}

// AsInt8Slice tries to convert data into a slice of int8
func AsInt8Slice(v interface{}) ([]int8, error) {
	switch w := v.(type) {
	case []int8:
		// ready
		return w, nil
	case []interface{}:
		// slice of something else, convert them to Int8
		out := make([]int8, 0, len(w))
		for _, o := range w {
			if n, err := AsInt8(o); err == nil {
				out = append(out, n)
			} else {
				// failed
				return nil, err
			}
		}
		return out, nil
	case interface{}:
		// promote single element to slice
		if n, err := AsInt8(v); err != nil {
			return nil, err
		} else {
			return []int8{n}, nil
		}
	default:
		return nil, InvalidTypeError(v)
	}
}

// AsUintSlice tries to convert data into a slice of uint
func AsUintSlice(v interface{}) ([]uint, error) {
	switch w := v.(type) {
	case []uint:
		// ready
		return w, nil
	case []interface{}:
		// slice of something else, convert them to Uint
		out := make([]uint, 0, len(w))
		for _, o := range w {
			if n, err := AsUint(o); err == nil {
				out = append(out, n)
			} else {
				// failed
				return nil, err
			}
		}
		return out, nil
	case interface{}:
		// promote single element to slice
		if n, err := AsUint(v); err != nil {
			return nil, err
		} else {
			return []uint{n}, nil
		}
	default:
		return nil, InvalidTypeError(v)
	}
}

// AsUint64Slice tries to convert data into a slice of uint64
func AsUint64Slice(v interface{}) ([]uint64, error) {
	switch w := v.(type) {
	case []uint64:
		// ready
		return w, nil
	case []interface{}:
		// slice of something else, convert them to Uint64
		out := make([]uint64, 0, len(w))
		for _, o := range w {
			if n, err := AsUint64(o); err == nil {
				out = append(out, n)
			} else {
				// failed
				return nil, err
			}
		}
		return out, nil
	case interface{}:
		// promote single element to slice
		if n, err := AsUint64(v); err != nil {
			return nil, err
		} else {
			return []uint64{n}, nil
		}
	default:
		return nil, InvalidTypeError(v)
	}
}

// AsUint32Slice tries to convert data into a slice of uint32
func AsUint32Slice(v interface{}) ([]uint32, error) {
	switch w := v.(type) {
	case []uint32:
		// ready
		return w, nil
	case []interface{}:
		// slice of something else, convert them to Uint32
		out := make([]uint32, 0, len(w))
		for _, o := range w {
			if n, err := AsUint32(o); err == nil {
				out = append(out, n)
			} else {
				// failed
				return nil, err
			}
		}
		return out, nil
	case interface{}:
		// promote single element to slice
		if n, err := AsUint32(v); err != nil {
			return nil, err
		} else {
			return []uint32{n}, nil
		}
	default:
		return nil, InvalidTypeError(v)
	}
}

// AsUint16Slice tries to convert data into a slice of uint16
func AsUint16Slice(v interface{}) ([]uint16, error) {
	switch w := v.(type) {
	case []uint16:
		// ready
		return w, nil
	case []interface{}:
		// slice of something else, convert them to Uint16
		out := make([]uint16, 0, len(w))
		for _, o := range w {
			if n, err := AsUint16(o); err == nil {
				out = append(out, n)
			} else {
				// failed
				return nil, err
			}
		}
		return out, nil
	case interface{}:
		// promote single element to slice
		if n, err := AsUint16(v); err != nil {
			return nil, err
		} else {
			return []uint16{n}, nil
		}
	default:
		return nil, InvalidTypeError(v)
	}
}

// AsUint8Slice tries to convert data into a slice of uint8
func AsUint8Slice(v interface{}) ([]uint8, error) {
	switch w := v.(type) {
	case []uint8:
		// ready
		return w, nil
	case []interface{}:
		// slice of something else, convert them to Uint8
		out := make([]uint8, 0, len(w))
		for _, o := range w {
			if n, err := AsUint8(o); err == nil {
				out = append(out, n)
			} else {
				// failed
				return nil, err
			}
		}
		return out, nil
	case interface{}:
		// promote single element to slice
		if n, err := AsUint8(v); err != nil {
			return nil, err
		} else {
			return []uint8{n}, nil
		}
	default:
		return nil, InvalidTypeError(v)
	}
}
