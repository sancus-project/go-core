package errors

import (
	"fmt"
)

type Wrapped struct {
	err error
	msg string
}

func (w *Wrapped) Error() string {
	var s string
	if w != nil {
		s = fmt.Sprintf("%s: %s", w.msg, w.err)
	}
	return s
}

func (w *Wrapped) Unwrap() error {
	if w != nil {
		return w.err
	}
	return nil
}

func Wrap(err error, s string, args... interface{}) error {
	if err != nil {
		if len(s) > 0 {
			if len(args) > 0 {
				s = fmt.Sprintf(s, args...)
			}

			err = &Wrapped{
				err: err,
				msg: s,
			}
		}
	}
	return err
}
