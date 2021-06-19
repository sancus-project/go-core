package errors

type Validator interface {
	Ok() bool
	Error() string
	Errors() []error
}

func AsValidator(err error) (Validator, bool) {
	for err != nil {
		if v, ok := err.(Validator); ok {
			return v, true
		}

		err = Unwrap(err)
	}
	return nil, false
}
