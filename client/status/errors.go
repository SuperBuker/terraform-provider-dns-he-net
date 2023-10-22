package status

import "fmt"

// ErrNoAuth is an error returned when the user is not authenticated.
// It is used when dns.he.net returns that the client is not authenticated.
// It includes two wrapped errors, ErrAuth and ErrHeNet.
type ErrNoAuth struct{}

func (e *ErrNoAuth) Error() string {
	return "not authenticated"
}

func (e *ErrNoAuth) Unwrap() []error {
	return []error{
		&ErrHeNet{e.Error()},
	}
}

// ErrPartialAuth is an error that is returned when authentication is incomplete.
type ErrPartialAuth struct{}

func (e *ErrPartialAuth) Error() string {
	return "authentication not completed"
}

func (e *ErrPartialAuth) Unwrap() []error {
	return []error{
		&ErrNoAuth{},
		&ErrHeNet{e.Error()},
	}
}

// ErrAuthFailed is an error that is returned when authentication fails.
type ErrAuthFailed struct {
	error string
}

func (e *ErrAuthFailed) Error() string {
	return fmt.Sprintf("authentication failed: %s", e.error)
}

func (e *ErrAuthFailed) Unwrap() []error {
	return []error{
		&ErrNoAuth{},
		&ErrHeNet{e.error},
	}
}

// ErrMissingOTPAuth is an error returned when the user is not fully authenticated.
// It is used when dns.he.net returns that the client lacks OTP authentication.
// It includes two wrapped errors, ErrAuth and ErrHeNet.
type ErrMissingOTPAuth struct{}

func (e *ErrMissingOTPAuth) Error() string {
	return "missing OTP authentication"
}

func (e *ErrMissingOTPAuth) Unwrap() []error {
	return []error{
		&ErrPartialAuth{},
		&ErrHeNet{e.Error()},
	}
}

// ErrOTPAuthFailed is an error that is returned when authentication fails.
type ErrOTPAuthFailed struct {
	error string
}

func (e *ErrOTPAuthFailed) Error() string {
	return fmt.Sprintf("authentication failed: %s", e.error)
}

func (e *ErrOTPAuthFailed) Unwrap() []error {
	return []error{
		&ErrAuthFailed{e.error},
		&ErrHeNet{e.error},
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

func errorScore(err error) int {
	switch err.(type) {
	case *ErrOTPAuthFailed:
		return 7
	case *ErrAuthFailed:
		return 6
	case *ErrMissingOTPAuth:
		return 5
	case *ErrPartialAuth:
		return 4
	case *ErrNoAuth:
		return 3
	case *ErrUnknownAuth:
		return 2
	case *ErrHeNet:
		return 1
	default:
		return 0
	}
}
