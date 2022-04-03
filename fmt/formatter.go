package fmt

import (
	"fmt"
)

func PassThrough(x interface{}, f fmt.State, c rune) {
	// compose format
	s := "%"
	// each possible ascii flag
	for i := 0; i < 128; i++ {
		if f.Flag(i) {
			s += string(rune(i))
		}
	}
	// width
	if wid, ok := f.Width(); ok {
		s += fmt.Sprintf("%d", wid)
	}
	// precision
	if prec, ok := f.Precision(); ok {
		s += fmt.Sprintf(".%d", prec)
	}
	// type
	s += string(c)

	fmt.Fprintf(f, s, x)
}
