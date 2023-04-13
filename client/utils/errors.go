package utils

type ErrNotImplemented struct {
}

func (e *ErrNotImplemented) Error() string {
	return "feature is not implemented"
}

func (e *ErrNotImplemented) Unwrap() []error {
	return []error{}
}
