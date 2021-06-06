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

func StackTrace(err error) Stack {
	for err != nil {
		if e, ok := err.(StackTracer); ok {
			st := e.StackTrace()
			if st != nil {
				return st
			}
		}

		err = Unwrap(err)
	}

	return nil
}
