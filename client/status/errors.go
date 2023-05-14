package status

// ErrAuthFailed is an error that is returned when authentication fails.
type ErrAuthFailed struct{}

func (e *ErrAuthFailed) Error() string {
	return "authentication failed"
}

// ErrNoAuth is an error returned when the user is not authenticated.
// It is used when dns.he.net returns that the client is not authenticated.
// It includes two wrapped errors, ErrAuth and ErrHeNet.
type ErrNoAuth struct{}

func (e *ErrNoAuth) Error() string {
	return "not authenticated"
}

func (e *ErrNoAuth) Unwrap() []error {
	return []error{
		&ErrAuthFailed{},
		&ErrHeNet{e.Error()},
	}
}

// ErrOTPAuth is an error returned when the user is not fully authenticated.
// It is used when dns.he.net returns that the client lacks OTP authentication.
// It includes two wrapped errors, ErrAuth and ErrHeNet.
type ErrOTPAuth struct{}

func (e *ErrOTPAuth) Error() string {
	return "missing OTP authentication"
}

func (e *ErrOTPAuth) Unwrap() []error {
	return []error{
		&ErrAuthFailed{},
		&ErrHeNet{e.Error()},
	}
}

// ErrUnknownAuth is an error returned when it's not possible to determine the
// authentication status.
// It includes two wrapped errors, ErrAuth and ErrHeNet.
type ErrUnknownAuth struct{}

func (e *ErrUnknownAuth) Error() string {
	return "unknown authentication status"
}

func (e *ErrUnknownAuth) Unwrap() []error {
	return []error{
		&ErrAuthFailed{},
		&ErrHeNet{e.Error()},
	}
}

// ErrHeNet is an error returned when dns.he.net returns an error.
// It contains the error message returned in the HTML response.
type ErrHeNet struct {
	error string
}

func (e *ErrHeNet) Error() string {
	return e.error
}
