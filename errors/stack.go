package errors

type StackTracer interface {
	StackTrace() Stack
}

type StackError struct {
	stack Stack
	err   error
}

func (s *StackError) Error() string {
	if s.err != nil {
		return s.err.Error()
	} else {
		return "stacktrace"
	}
}

func (s *StackError) StackTrace() Stack {
	return s.stack
}

func (s *StackError) Unwrap() error {
	return s.err
}

func WithStackTrace(skip int, err error) error {
	return &StackError{
		err:   err,
		stack: BackTrace(1 + skip),
	}
}

func NewWithStackTrace(skip int, s string, args ...interface{}) error {
	return &StackError{
		err:   New(s, args...),
		stack: BackTrace(1 + skip),
	}
}

func AsStackTracer(err error) (StackTracer, bool) {
	for err != nil {
		if e, ok := err.(StackTracer); ok {
			st := e.StackTrace()
			if len(st) > 0 {
				return e, true
			}
		}

		err = Unwrap(err)
	}

	return nil, false
}

func StackTrace(err error) Stack {

	if e, ok := AsStackTracer(err); ok {
		return e.StackTrace()
	}

	return nil
}
