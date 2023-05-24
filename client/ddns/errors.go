package ddns

import (
	"fmt"
)

// ErrDDNS error is returned when interation against the DDNS endpoint fails.
type ErrDDNS struct {
	error string
}

func (e *ErrDDNS) Error() string {
	return fmt.Sprintf("ddns update failed: %q", e.error)
}

func (e *ErrDDNS) Unwrap() []error {
	return []error{}
}

// ErrAuthFailed is an error that is returned when authentication fails.
type ErrAuthFailed struct{}

func (e *ErrAuthFailed) Error() string {
	return "authentication failed"
}

func (e *ErrAuthFailed) Unwrap() []error {
	return []error{
		&ErrDDNS{"badauth"},
	}
}

// ErrAbuse is an error that is returned when too many requests were performed.
type ErrAbuse struct {
}

func (e *ErrAbuse) Error() string {
	return "authentication blocked due to abuse"
}

func (e *ErrAbuse) Unwrap() []error {
	return []error{
		&ErrDDNS{"abuse"},
	}
}

// ErrField an error that is returned when the request is malformed.
type ErrField struct {
	error string
}

func (e *ErrField) field() (field string, ok bool) {
	field, ok = map[string]string{
		"noipv4": "myip",
		"noipv6": "myip",
		"badip":  "myip",
		"notxt":  "txt",
	}[e.error]

	return
}

func (e *ErrField) Error() string {
	if field, ok := e.field(); ok {
		return fmt.Sprintf("missing or malformed field: %q", field)
	}

	return fmt.Sprintf("malformed request: %q", e.error)
}

func (e *ErrField) Unwrap() []error {
	return []error{
		&ErrDDNS{e.error},
	}
}

// ErrAPI is an error that is returned when the http client returns an error.
type ErrAPI struct {
	error error
}

func (e *ErrAPI) Error() string {
	return e.error.Error()
}

func (e *ErrAPI) Unwrap() []error {
	return []error{
		e.error,
		&ErrDDNS{e.error.Error()},
	}
}

// ErrUnknown is an error that is returned when the error cause is unknown.
type ErrUnknown struct {
	error string
}

func (e *ErrUnknown) Error() string {
	return fmt.Sprintf("unknown error: %q", e.error)
}

func (e *ErrUnknown) Unwrap() []error {
	return []error{
		&ErrDDNS{e.error},
	}
}
