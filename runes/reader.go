package runes

import (
	"bytes"
	"io"
	"strings"

	"github.com/SteelSeries/bufrr"
)

type (
	Reader      = bufrr.Reader
	RunePeeker  = bufrr.RunePeeker
	RuneReader  = io.RuneReader
	RuneScanner = io.RuneScanner
)

func NewReader(in io.Reader) *Reader {
	return bufrr.NewReader(in)
}

func NewReaderBytes(b []byte) *Reader {
	return NewReader(bytes.NewReader(b))
}

func NewReaderString(s string) *Reader {
	return NewReader(strings.NewReader(s))
}

func NewReaderSize(in io.Reader, size int) *Reader {
	return bufrr.NewReaderSize(in, size)
}

func NewReaderBytesSize(b []byte, size int) *Reader {
	return NewReaderSize(bytes.NewReader(b), size)
}

func NewReaderStringSize(s string, size int) *Reader {
	return NewReaderSize(strings.NewReader(s), size)
}
