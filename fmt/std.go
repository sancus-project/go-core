package fmt

import (
	"fmt"
	"io"
)

type (
	Formatter  = fmt.Formatter
	GoStringer = fmt.GoStringer
	State      = fmt.State
	Stringer   = fmt.Stringer
)

func Errorf(format string, args ...interface{}) error {
	return fmt.Errorf(format, args...)
}

func Fprint(w io.Writer, args ...interface{}) (int, error) {
	return fmt.Fprint(w, args...)
}

func Fprintln(w io.Writer, args ...interface{}) (int, error) {
	return fmt.Fprintln(w, args...)
}

func Fprintf(w io.Writer, format string, args ...interface{}) (int, error) {
	return fmt.Fprintf(w, format, args...)
}

func Print(args ...interface{}) (int, error) {
	return fmt.Print(args...)
}

func Println(args ...interface{}) (int, error) {
	return fmt.Println(args...)
}

func Printf(format string, args ...interface{}) (int, error) {
	return fmt.Printf(format, args...)
}

func Sprint(args ...interface{}) string {
	return fmt.Sprint(args...)
}

func Sprintln(args ...interface{}) string {
	return fmt.Sprintln(args...)
}

func Sprintf(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}
