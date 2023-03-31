package status

type ErrAuth struct{}

func (e *ErrAuth) Error() string {
	return "generic auhentication error"
}

type ErrNoAuth struct{}

func (e *ErrNoAuth) Error() string {
	return "not authenticated"
}

func (e *ErrNoAuth) Unwrap() []error {
	return []error{
		&ErrAuth{},
		&ErrHeNet{e.Error()},
	}
}

type ErrOTPAuth struct{}

func (e *ErrOTPAuth) Error() string {
	return "missing OTP authentication"
}

func (e *ErrOTPAuth) Unwrap() []error {
	return []error{
		&ErrAuth{},
		&ErrHeNet{e.Error()},
	}
}

type ErrUnknownAuth struct{}

func (e *ErrUnknownAuth) Error() string {
	return "unknown authentication status"
}

func (e *ErrUnknownAuth) Unwrap() []error {
	return []error{
		&ErrAuth{},
		&ErrHeNet{e.Error()},
	}
}

type ErrHeNet struct {
	error string
}

func (e *ErrHeNet) Error() string {
	return e.error
}
